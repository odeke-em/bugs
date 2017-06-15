package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	var targetURL string
	flag.StringVar(&targetURL, "url", "https://localhost:9999/", "the url to invoke")
	flag.Parse()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http2.ConfigureTransport(tr)
	client := &http.Client{Transport: tr}
	res, err := client.Get(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	statusOK := res.StatusCode >= 200 && res.StatusCode <= 299
	if !statusOK {
		log.Fatalf("bad status: %q", res.Status)
	}
}
