package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"sync"
	"time"
)

func main() {
	log.SetFlags(0)

	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		blob, _ := ioutil.ReadAll(io.LimitReader(r.Body, 100))
		log.Printf("Server: %s", blob)
		w.Write([]byte("Hello, world!"))
		w.(http.Flusher).Flush()
		r.Body.Close()
	}))
	defer cst.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	prc, pwc := io.Pipe()

	var triggerSendHeadersOnce sync.Once
	sendHeaders := make(chan bool)

	go func() {
		defer pwc.Close()

		<-sendHeaders

		for i := 0; i < 100; i++ {
			select {
			case <-time.After(450 * time.Millisecond):
				pwc.Write([]byte("Hello from client"))
			case <-ctx.Done():
				return
			}
		}
	}()

	trace := &httptrace.ClientTrace{
		WroteHeaders: func() {
			triggerSendHeadersOnce.Do(func() {
				log.Println("Closing sendHeaders")
				close(sendHeaders)
			})
		},
		WroteHeaderField: func(key string, value []string) {
			log.Printf("Header %q: %#v\n", key, value)
		},
		WroteRequest: func(ri httptrace.WroteRequestInfo) {
			log.Printf("WroteRequestInfo: %#v", ri)
		},
	}

	req, _ := http.NewRequest("POST", cst.URL, prc)
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	blob, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	log.Printf("Client: %s", blob)
}
