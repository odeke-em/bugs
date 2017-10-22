package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	var certFile, keyFile string

	flag.StringVar(&certFile, "cert", "cert.pem", "the certificate")
	flag.StringVar(&keyFile, "key", "key.pem", "the key file")
	flag.IntVar(&port, "port", 8877, "the port to run on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServeTLS(addr, certFile, keyFile, &handler{}); err != nil {
		log.Fatal(err)
	}
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hdr := w.Header()
	if r.URL.Path == "/" {
		hdr.Set("Content-Type", "text/plain")
		hdr.Set("Age", "3960")
		hdr.Set("Connection", "keep-alive")
		hdr.Set("Date", "keep-alive")
		hdr.Set("Server", "Netlify")
		hdr.Set("Cache-Control", "public, max-age=0, must-revalidate")
		hdr.Set("Date", "Sun, 22 Oct 2017 04:32:22 GMT")
		if false && r.URL.RawPath == "" { // Can't seem to differentiate between localhost:8877 & localhost:8877/
			hdr.Set("Location", "/")
		} else {
			hdr.Set("Location", "/blog/dont-design-emails/")
		}
		w.WriteHeader(301)
	} else {
		hdr.Set("Cache-Control", "public, max-age=0, must-revalidate")
		hdr.Set("Content-Type", "text/html; charset=UTF-8")
		hdr.Set("Date", "Sun, 22 Oct 2017 04:32:23 GMT")
		hdr.Set("Age", "3960")
		hdr.Set("Connection", "keep-alive")
		hdr.Set("Server", "Netlify")
		fmt.Fprintf(w, "Hello Gopher")
	}
	w.(http.Flusher).Flush()
}
