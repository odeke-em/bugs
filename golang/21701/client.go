package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	var theURL string
	flag.StringVar(&theURL, "url", "http://localhost:8877", "the server address")
	flag.Parse()

	proxyURL, _ := url.Parse(theURL)
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,
			DisableKeepAlives:   true,
		},
	}
	r, _ := http.NewRequest("POST", "https://golang.org/", nil)
	res, err := client.Do(r)
	if err != nil {
		log.Fatalf("res: %v", err)
	}
	wire, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wire: %s\n", wire)
}
