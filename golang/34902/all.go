package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
)

func main() {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this call was relayed by the reverse proxy")
	}))
	defer cst.Close()

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to find available port: %v", err)
	}
	defer ln.Close()

	backendURL, _ := url.Parse(cst.URL)
	rpxy := httputil.NewSingleHostReverseProxy(backendURL)
	addr := ln.Addr().String()

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		if err := http.Serve(ln, rpxy); err != nil {
			log.Fatalf("Reverse proxy serve error: %v", err)
		}
	}()

	runClients(20, addr, &wg)
}

func runClients(n int, addr string, wg *sync.WaitGroup) {
	fullURL := "http://" + addr

	// Request needs a larger body to see the error happen more quickly
	// It is reproducable with smaller body sizes, but it takes longer to fail
	bodyString := strings.Repeat("a", 2048)
	client := &http.Client{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				buf := bytes.NewBufferString(bodyString)
				req, _ := http.NewRequest("POST", fullURL, buf)
				req.Header.Add("Expect", "100-continue")

				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("Request Failed: %s\n", err.Error())
				} else {
					resp.Body.Close()
				}
			}
		}()
	}
}
