package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func receive(rw http.ResponseWriter, req *http.Request) {
	log.Printf("req: %#v", req.Header)
	slurp, _ := ioutil.ReadAll(req.Body)
	log.Printf("slurp: %s", slurp)

	switch req.Method {
	case "GET":
		if len(slurp) > 0 {
			http.Error(rw, fmt.Sprintf("got %q", slurp), http.StatusForbidden)
			return
		}

		ctLen := req.ContentLength
		if ctLen >= 1 {
			http.Error(rw, "not expecting a body for GET", http.StatusForbidden)
			return
		}
	}

	rw.Write([]byte("OK Gopher"))
}

func main() {
	var port int
	var theURL string
	var asServer bool
	var http2 bool

	flag.BoolVar(&asServer, "as-server", false, "run it as a server")
	flag.BoolVar(&http2, "http2", false, "run it as an http2 server")
	flag.IntVar(&port, "port", 9898, "port to run it on")
	flag.StringVar(&theURL, "url", "https://localhost:9898", "the URL to fetch")
	flag.Parse()

	if asServer {
		http.HandleFunc("/", receive)
		addr := fmt.Sprintf(":%d", port)
		log.Printf("running on %s", addr)
		if http2 {
			if err := http.ListenAndServeTLS(addr, "./cert.pem", "./key.pem", nil); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := http.ListenAndServe(addr, nil); err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Printf("the URL: %s", theURL)

	c := http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	// According to the CloudFront documentation for a request behavior, if the
	// request is GET and includes a body, it returns a 403 Forbidden. See the
	// documentation here:
	// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#RequestCustom-get-body

	var body bytes.Buffer

	r, err := http.NewRequest("GET", theURL, ioutil.NopCloser(&body))
	if err != nil {
		log.Fatalf("error creating the request: %s", err)
	}
	if false {
		blob, _ := httputil.DumpRequestOut(r, true)
		log.Printf("dumpRequest: %q", blob)
	}

	res, err := c.Do(r)
	if err != nil {
		log.Fatalf("error doing the request: %s", err)
	}
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	log.Printf("response status for an HTTP/2 request: %s", res.Status)

	// doing the same request without HTTP/2 does work
	tr := c.Transport.(*http.Transport)
	tr.TLSNextProto = make(map[string]func(string, *tls.Conn) http.RoundTripper)
	c.Transport = tr

	r, err = http.NewRequest("GET", theURL, ioutil.NopCloser(&body))
	if err != nil {
		log.Fatalf("error creating the request: %s", err)
	}
	if false {
		blob, _ := httputil.DumpRequestOut(r, true)
		log.Printf("dumpRequest: %q", blob)
	}

	res, err = c.Do(r)
	if err != nil {
		log.Fatalf("error doing the request: %s", err)
	}
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	log.Printf("response status for an HTTP/1 request: %s", res.Status)
}
