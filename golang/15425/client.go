package main

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/http2"
)

func main() {
	serverURL := "https://localhost:1333"
	if envServURL := os.Getenv("SERVER_URL"); envServURL != "" {
		serverURL = envServURL
	}

	certs, err := tls.LoadX509KeyPair("key.crt", "key.key")
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{certs},
			InsecureSkipVerify: true,
		},
	}
	http2.ConfigureTransport(tr)

	client := &http.Client{
		Transport: tr,
	}

	// Make a GigaByte of data
	raw := strings.Repeat("A", 1024*1024*1024)
	log.Printf("len(data)=%v\n", len(raw))
	sr := strings.NewReader(raw)
	req, err := http.NewRequest("POST", serverURL, sr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("client.Do next\n")
	res, err := client.Do(req)
	log.Printf("client.Do done\n")
	if err != nil {
		log.Fatal(err)
	}
	n, err := io.Copy(ioutil.Discard, res.Body)
	log.Printf("done Reading the server response's body n=(%v) err=%v\nres.headers: %v\n", n, err, res.Header)
	_ = res.Body.Close()
}
