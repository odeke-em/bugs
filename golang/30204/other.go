package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"golang.org/x/net/http2"
)

func main() {
	cst := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqBlob, _ := httputil.DumpRequest(req, true)
		rw.Write(reqBlob)
	}))
        http2.ConfigureServer(cst.Config, nil)
        cst.StartTLS()
	defer cst.Close()

        log.Printf("URL: %s\n", cst.URL)

	req, _ := http.NewRequest("GET", cst.URL, nil)
	// req.Header.Set("Upgrade", "websocket")
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	http2.ConfigureTransport(tr)
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch: %v\n", err)
	}
	blob, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	log.Printf("\nResponse from server:\n\n%s\n", blob)
}
