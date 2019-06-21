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
	"sync/atomic"
	"syscall"
	"text/template"
	"time"
)

const maxIdleTimeout = 6 * time.Second

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to find an available port: %v", err)
	}
	defer ln.Close()

	addr := ln.Addr().String()
	url := "http://" + addr

	mux := http.NewServeMux()

	var reqCounter int64
	onLag := make(chan bool)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from: %v\n\n", r.RemoteAddr)

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
		log.Printf("Now going to lag for %s", lagTime)
		time.Sleep(lagTime)
		conn.(*net.TCPConn).SetLinger(0)
		conn.Close()
	})

	log.Printf("Listening at %q", addr)

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
	// process in pause at:
	//              maxIdleTimeout/3
	// We'll then revive the process at:
	//              maxIdleTimeout * 4/5
	// way after the connection has been closed so that from the perspective of
	// the client it sees a connection that hasn't yet expired and isn't idle
	// but the server would have already terminated it, thus triggering ECONNRESET.

	stdoutStderr := os.Stdout
	cmd, done, err := runClient(url, stdoutStderr, stdoutStderr)
	if err != nil {
		log.Printf("Failed to run client command: %v", err)
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
	log.Printf("Pausing for %s before reviving client with pid: %d\n", timeBeforeContinue, clientProc.Pid)
	<-time.After(timeBeforeContinue)
	// Revive the client.
	clientProc.Signal(syscall.SIGCONT)

	// The client should see an ECONNRESET due to a now stale connection
	// and it should have provided a now reused connection.
	<-onLag
	// Wait for:
	//      maxIdleTimeout * 2/3

	<-time.After(5 * maxIdleTimeout / 3)
	// Finally kill the client process.
	clientProc.Signal(syscall.SIGKILL)

	_ = cmd.Wait()
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
        cmd.Env =  append(os.Environ(), "GODEBUG=netdns=1")
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("Failed to start the client binary: %v", err)
	}

	return
}

var tmpl *template.Template

func init() {
	log.SetPrefix("server: ")
	blob, err := ioutil.ReadFile("./client.go")
	if err != nil {
		log.Fatalf("Failed to read the source code for client.go: %v", err)
	}
	tmpl = template.Must(template.New("client").Parse(string(blob)))
}
