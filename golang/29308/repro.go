package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"syscall"
	"text/template"
	"time"
)

const maxIdleTimeout = 3 * time.Second

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to find an available port: %v", err)
	}
	defer ln.Close()

	fmt.Printf("This test could take about %s (or more) to run\n", maxIdleTimeout*9)

	addr := ln.Addr().String()
	url := "http://" + addr

	mux := http.NewServeMux()
	serverBuf := new(bytes.Buffer)
	serverLog := log.New(serverBuf, "server:", 0)

	var reqCounter int64
	onLag := make(chan bool)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		serverLog.Printf("Request from: %v\n\n", r.RemoteAddr)

		n := atomic.AddInt64(&reqCounter, 1)
		switch n {
		case 1:
			w.Write([]byte("Server responding ASAP"))
			return
		}

		// An even connection count means that the previous one is being reused.
		// This route is meant invoke conn.Close() to produce an ECONNRESET that'll
		// then trigger the malfesence on the client's Transport
		// as reported in https://golang.org/issues/29308
		conn, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		onLag <- true
		// Sleep for sometime and then close the connection.
		lagTime := maxIdleTimeout / 2
		serverLog.Printf("Now going to lag for %s", lagTime)
		time.Sleep(lagTime)
		conn.(*net.TCPConn).SetLinger(0)
		conn.Close()
	})

	serverLog.Printf("Listening at %q", addr)

	go func() {
		// Not doing:
		//      log.Fatal(http.Serve(ln, mux))
		// because when we invoke ln.Close() in the defer,
		// it'll also cause http.Serve to fail with:
		//      "accept tcp [::]:<port>: use of closed network connection"
		// which is just a spurious error.
		_ = http.Serve(ln, mux)
	}()

	// Spin up the client, let it make a request to the server above,
	// and while the connection is still alive, we'll put the client's
	// process in pause and then revive it at:
	//              maxIdleTimeout/3
	// We'll then revive the process at:
	//              maxIdleTimeout * 4/5
	// way after the connection has been closed so that from the perspective of
	// the client it sees a fresh connection that hasn't yet been reused nor expired
	// but the server would have already terminated it, thus triggering ECONNRESET.

	stdoutStderr := new(bytes.Buffer)
        ssr := io.MultiWriter(os.Stdout, stdoutStderr)
	cmd, done, err := runClient(url, ssr, ssr)
	if err != nil {
		serverLog.Printf("Failed to run client command: %v", err)
		return
	}
	defer done()

	// Wait until the client has hit the <-onLag point.
	<-onLag

	clientProc := cmd.Process
	// Now stop the client process.
	clientProc.Signal(syscall.SIGTSTP)

	// Wait for about:
	//      maxIdleTimeout * 1/3
	timeBeforeContinue := maxIdleTimeout / 3
	serverLog.Printf("Pausing for %s before reviving client with pid: %d\n", timeBeforeContinue, clientProc.Pid)
	<-time.After(timeBeforeContinue * 3 * 5)
	// Revive the client.
	clientProc.Signal(syscall.SIGCONT)

	// The client should see an ECONNRESET due to a now stale connection
	// and it should have provided a now reused connection.
	<-onLag
	<-time.After(timeBeforeContinue / 6)

	// Since we now have the fresh connection, let's also pause
	// and wait for ECONNRESET to be triggered.
	clientProc.Signal(syscall.SIGTSTP)
	<-time.After(maxIdleTimeout)

	// Now revive the client process back up.
	clientProc.Signal(syscall.SIGCONT)

	<-time.After(5 * maxIdleTimeout / 3)

	// Finally kill the client process.
	clientProc.Signal(syscall.SIGKILL)

	_ = cmd.Wait()

	clientLogs := stdoutStderr.String()
	badStrings := []string{"Failed to make request", "read: connection reset by peer"}
	failed := false
	for _, badString := range badStrings {
		if strings.Contains(clientLogs, badString) {
			fmt.Printf("Unfortunately contains: %q\n", badString)
			failed = true
		}
	}

	if failed {
		fmt.Printf("Client log:\n\n%s", clientLogs)
	} else {
		fmt.Println("Congrats! Test doesn't seem to suffered from connection resets")
	}
}

func runClient(url string, stdout, stderr io.Writer) (cmd *exec.Cmd, done func(), err error) {
	tmpDir, err := ioutil.TempDir("", "client")
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create tmpDir: %v", err)
	}
	defer func() {
		if err == nil {
			done = func() { os.RemoveAll(tmpDir) }
		} else {
			os.RemoveAll(tmpDir)
			done = nil
		}
	}()

	buf := new(bytes.Buffer)
	params := struct {
		MaxIdleTimeout string
		URL            string
	}{
		MaxIdleTimeout: maxIdleTimeout.String(),
		URL:            url,
	}
	if err := tmpl.Execute(buf, params); err != nil {
		return nil, nil, fmt.Errorf("Failed to create client source code: %v", err)
	}

	clientMainGoPath := filepath.Join(tmpDir, "main.go")
	if err := ioutil.WriteFile(clientMainGoPath, buf.Bytes(), 0644); err != nil {
		return nil, nil, fmt.Errorf("Failed to write client main.go source code to file: %v", err)
	}
	clientBinaryPath := filepath.Join(tmpDir, "binary")

	// Then build it as a binary
	out, err := exec.Command("go", "build", "-o", clientBinaryPath, clientMainGoPath).CombinedOutput()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to build client binary: %v\n%s", err, out)
	}

	// Once we have the binary, we just need to run it.
	cmd = exec.Command(clientBinaryPath)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Env = append(os.Environ(), "GODEBUG=netdns=1")
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("Failed to start the client binary: %v", err)
	}

	return
}

var tmpl = template.Must(template.New("client").Parse(clientMainGo))

const clientMainGo = `
package main

// This file is just a template that'll get populated
// with values set by the server. After population, this
// code shall be run as a child process and also controlled
// by various signals sent to it, in order to mimick the
// conditions from https://golang.org/issues/29308

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "{{.URL}}"

var maxIdleTimeout time.Duration

func init() {
	log.SetPrefix("client: ")
	idleTimeout, err := time.ParseDuration("{{.MaxIdleTimeout}}")
	if err != nil {
		log.Fatalf("Failed to parse MaxIdleTimeout: %v", err)
	}
	maxIdleTimeout = idleTimeout
}

func main() {
	tr := &http.Transport{
		IdleConnTimeout: maxIdleTimeout,
	}
	client := &http.Client{
		Transport: tr,
	}

	path := "/hello"
	for {
		purl := url + path
		log.Printf("URL: %s\n", purl)
		res, err := client.Get(purl)
		if err != nil {
			log.Fatalf("Failed to make request(%q): %v", purl, err)
		}
		blob, err := ioutil.ReadAll(res.Body)
		log.Printf("Blob: %s\n", blob)
		if err != nil {
			log.Fatalf("%q failed to read body: %v", purl, err)
		}
		pauseTime := maxIdleTimeout / 4
		log.Printf("Pausing for %s\n\n", pauseTime)
		<-time.After(pauseTime)
	}
}
`
