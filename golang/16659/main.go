package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func main() {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "2000")
		fmt.Fprintln(w, "this call was relayed by the reverse proxy")
	}))
	defer backendServer.Close()

	rpURL, err := url.Parse(backendServer.URL)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	rproxy := httputil.NewSingleHostReverseProxy(rpURL)
	rproxy.ErrorLog = log.New(&buf, "18: ", log.Lshortfile)
	frontendProxy := httptest.NewServer(rproxy)

	defer frontendProxy.Close()

	resp, err := http.Get(frontendProxy.URL)
	if err != nil {
		log.Fatalf("getErr: %v", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("buff: %s\n", buf.Bytes())
	if err != nil {
		log.Fatalf("slurpErr: %v", err)
	}

	fmt.Printf("b %s", b)
}
