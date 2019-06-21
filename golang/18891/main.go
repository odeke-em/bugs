package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	theURL := flag.String("url", "http://localhost:8888/", "the URL to invoke")
	flag.Parse()
	c := http.Client{}

	// According to the CloudFront documentation for a request behavior, if the
	// request is GET and includes a body, it returns a 403 Forbidden. See the
	// documentation here:
	// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#RequestCustom-get-body
	var body io.Reader = func() io.Reader {
		buf := new(bytes.Buffer)
		buf.Write([]byte("a"))
		return buf
	}()
	body = ioutil.NopCloser(new(bytes.Buffer))
	r, err := http.NewRequest("GET", *theURL, body)
	if err != nil {
		log.Fatalf("error creating the request: %s", err)
	}

	res, err := c.Do(r)
	if err != nil {
		log.Fatalf("error doing the request: %s", err)
	}
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	log.Printf("response status for an HTTP/2 request: %s", res.Status)

	// doing the same request without HTTP/2 does work
	c.Transport = &http.Transport{
		TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
	}
	r, err = http.NewRequest("GET", *theURL, nil)
	if err != nil {
		log.Fatalf("error creating the request: %s", err)
	}

	res, err = c.Do(r)
	if err != nil {
		log.Fatalf("error doing the request: %s", err)
	}
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	log.Printf("response status for an HTTP/1 request: %s", res.Status)
}
