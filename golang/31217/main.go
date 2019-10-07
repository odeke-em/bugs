package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS12,
			},
		},
	}
	res, err := hc.Get("https://www.google.com")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	blob, _ := ioutil.ReadAll(res.Body)
	println(string(blob))
}

func foo() {
	conn, err := tls.Dial("tcp", "googleapis.com:443", &tls.Config{
		// TODO: please properly load your TLSClient configuration here.
		// Just using InsecureSkipVerify for easy testing.
		InsecureSkipVerify: true,
		MaxVersion:         tls.VersionTLS12,
	})
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer conn.Close()

	// Write something to the connection.
	conn.Write([]byte("POST / HTTP/1.1\r\nPath: /upload/\r\n\r\n"))
}
