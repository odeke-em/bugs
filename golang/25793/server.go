package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

const NumOfRequests = 20

func main() {
	log.SetFlags(0)
	canDebug := os.Getenv("CANDEBUG") != ""
	cst := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 10ms processing time
		if canDebug {
			dump, _ := httputil.DumpRequest(r, true)
			log.Printf("\nServer: \033[32m%s\033[00m\n", dump)
			time.Sleep(10 * time.Millisecond)
		}
		io.WriteString(w, "ok")
	}))
	defer cst.Close()

	// Give the server some time to setup
	time.Sleep(100 * time.Millisecond)

	// Set up the client, with HTTP/2.0 support, but keepalives disabled
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy: func(req *http.Request) (*url.URL, error) {
			// Alter the request context and return no proxy
			*req = *req.WithContext(context.WithValue(req.Context(), "Proxy", 1))
			return nil, nil
		},
		DisableKeepAlives: true,
	}
	http2.ConfigureTransport(transport)
	client := &http.Client{Transport: transport}

	// Execute some requests to the same host, with a 1ms
	wg := sync.WaitGroup{}
	wg.Add(NumOfRequests)
	for i := 0; i < NumOfRequests; i++ {
		// 10 ms wait time between requests
		time.Sleep(10 * time.Millisecond)
		go func(id int) {
			defer wg.Done()
			req, err := http.NewRequest("GET", fmt.Sprintf("%s?i=%d", cst.URL, id), nil)
			handleErr(err)

			res, err := client.Do(req)
			handleErr(err)
			defer res.Body.Close()

			log.Printf("Client: \033[33m%v\033[00m", res.Request.Context())
		}(i)
	}
	wg.Wait()
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
