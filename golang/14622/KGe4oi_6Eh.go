package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s", req.RemoteAddr, req.Proto)
}

func main() {
	http.HandleFunc("/hello", HelloServer)

	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	cipherSuites := []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		25551,
	}

	tlsConfig := tls.Config{
		CipherSuites: cipherSuites,
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"h2"},
	}

	server := http.Server{
		TLSConfig: &tls.Config{},
	}

	ln, err := tls.Listen("tcp", ":54321", &tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(ln); err != nil {
		log.Fatal(err)
	}
}
