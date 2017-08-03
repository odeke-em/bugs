package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var http1 bool
	flag.BoolVar(&http1, "http1", false, "use http1 to invoke the server")
	flag.Parse()

	tr := &http.Transport{}
	if http1 {
		tr.TLSClientConfig = &tls.Config{
			NextProtos: []string{"h1"},
		}
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}
	resp, err := client.Get("https://www.southampton.ac.uk/")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
