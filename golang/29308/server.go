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
	"syscall"
	"text/template"
	"time"
)

const maxIdleTimeout = 10 * time.Second

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to find an available port: %v", err)
	}
	defer ln.Close()

	addr := ln.Addr().String()
	url := "http://" + addr

	mux := http.NewServeMux()
	mux.HandleFunc("/initial", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("There you go!\n"))
	})
	mux.HandleFunc("/lag", func(w http.ResponseWriter, r *http.Request) {
		conn, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Request from: %v\n", r.RemoteAddr)
		time.Sleep(maxIdleTimeout / 2)
		conn.(*net.TCPConn).SetLinger(0)
		conn.Close()
	})

	log.Printf("Listening at %q", addr)

	// revive it with SIGCONT and let it go on with its process.
	go func() {
		log.Fatal(http.Serve(ln, mux))
	}()

	// Per client, we'll make a request at maxIdleTimeout/2
	// then send to the process SIGTSTP, and at 4/3 maxIdleTimeout
	// Procedurally.
	// Let it run itself first
	stdoutStderr := new(bytes.Buffer)
	cmd, done, err := runClient(url, stdoutStderr, stdoutStderr)
	if err != nil {
		log.Printf("Failed to run client command: %v", err)
		return
	}
	defer done()

	clientProc := cmd.Process
	<-time.After(2 * maxIdleTimeout / 3)
	clientProc.Signal(syscall.SIGTSTP)
	<-time.After(2 + (maxIdleTimeout / 3))
	clientProc.Signal(syscall.SIGCONT)
	<-time.After(2 * maxIdleTimeout / 3)
	clientProc.Signal(syscall.SIGKILL)

	err = cmd.Wait()
        log.Printf("Waited till client terminated: err: %v\nOutput:\n%s\n\n", err, stdoutStderr.Bytes())
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
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("Failed to start the client binary: %v", err)
	}

	return
}

var tmpl *template.Template

func init() {
	blob, err := ioutil.ReadFile("./client.go")
	if err != nil {
		log.Fatalf("Failed to read the source code for client.go: %v", err)
	}
	tmpl = template.Must(template.New("client").Parse(string(blob)))
}
