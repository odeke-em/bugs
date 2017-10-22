package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/http2"
)

func main() {
	var certFile, keyFile string
	flag.StringVar(&certFile, "cert", "cert.pem", "the certificate")
	flag.StringVar(&keyFile, "key", "key.pem", "the key file")
	flag.Parse()

	badURL, err := url.Parse("https://localhost:8877")
	if err != nil {
		panic(err)
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("loadX509: %v", err)
	}
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	if err := http2.ConfigureTransport(tr); err != nil {
		log.Fatalf("http2.ConfigureTransport: %v", err)
	}
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("HEAD", badURL.String(), nil)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
