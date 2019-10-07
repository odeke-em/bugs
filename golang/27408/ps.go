package main

import (
	"net/http"
	"net"
	"sync"
	"io"
	"fmt"
	"bytes"
)

func echo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	var buf bytes.Buffer
	io.Copy(&buf, r.Body)
	resp := fmt.Sprintf("method: %s, body: %s", r.Method, buf.String())
	w.Write([]byte(resp))
}

type Proxy struct {
	Backend string
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Receive a request from %v\n", r.RemoteAddr)
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}

	backendConn, err := net.Dial("tcp", p.Backend)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer backendConn.Close()

	clientConn, bufrw, err := hj.Hijack()
	if err != nil {
		return
	}

	defer clientConn.Close()

	if err := r.Write(backendConn); err != nil {
		return
	}

	var wg = sync.WaitGroup{}
	wg.Add(2)
	copy := func(dst io.Writer, src io.Reader) {
		io.Copy(dst, src)
		if conn, ok := dst.(interface {
			CloseWrite() error
		}); ok {
			conn.CloseWrite()
		}
		wg.Done()
	}

	bufferedReader := io.LimitReader(bufrw, int64(bufrw.Reader.Buffered()))
	mr := io.MultiReader(bufferedReader, clientConn)

	go copy(clientConn, backendConn)
	go copy(backendConn, mr)
	wg.Wait()
}

func main() {
	go http.ListenAndServe(":8081", http.HandlerFunc(echo))

	p := &Proxy{Backend: "localhost:8081"}
	http.ListenAndServe(":8080", p)
}