package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"

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

	res, err := client.Get(serverURL)
	if err != nil {
		log.Fatal(err)
	}
	n, err := io.Copy(os.Stdout, res.Body)
	log.Printf("done Reading the server response's body n=(%v) err=%v\nres.headers: %v\n", n, err, res.Header)
	_ = res.Body.Close()
}
