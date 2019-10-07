package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	log.SetFlags(0)
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to grab an available address: %v", err)
	}
	defer ln.Close()

	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqBlob, _ := httputil.DumpRequest(req, true)
		rw.Write(reqBlob)
	})
	srv := &http.Server{Handler: handler}
	go func() {
		_ = srv.ServeTLS(ln, "cert.pem", "key.pem")
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())

	addr := "https://localhost:" + port
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	http2.ConfigureTransport(tr)
	client := &http.Client{Transport: tr}

	// Just perform a canary request to show what the
	// response looks like before attempting an upgrade.
	canaryReq, _ := http.NewRequest("GET", addr, nil)
	canaryRes, _ := client.Do(canaryReq)
	canaryResBlob, _ := httputil.DumpResponse(canaryRes, true)
	log.Printf("CanaryResponse:\n%s\n\n", canaryResBlob)
	_ = canaryRes.Body.Close()

	// Now finally the request that we'd like.
	req, _ := http.NewRequest("GET", addr, nil)
	req.Header.Set("Upgrade", "Websocket")
	req.Header.Set("Connection", "upgrade")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch: %v\n", err)
	}
	resBlob, _ := httputil.DumpResponse(res, true)
	log.Printf("Response from server:\n%s\n", resBlob)
	_ = res.Body.Close()

	for {
		<-time.After(100 * time.Second)
	}
}
