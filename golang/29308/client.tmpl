package main

// This file is just a template that'll get populated
// with values set by the server. After population, this
// code shall be run as a child process and also controlled
// by various signals sent to it, in order to mimick the
// conditions from https://golang.org/issues/29308

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "{{.URL}}"

var maxIdleTimeout time.Duration

func init() {
	log.SetPrefix("client: ")
	idleTimeout, err := time.ParseDuration("{{.MaxIdleTimeout}}")
	if err != nil {
		log.Fatalf("Failed to parse MaxIdleTimeout: %v", err)
	}
	maxIdleTimeout = idleTimeout
}

func main() {
	tr := &http.Transport{
		IdleConnTimeout: maxIdleTimeout,
	}
	client := &http.Client{
		Transport: tr,
	}

	path := "/hello"
	for {
		purl := url + path
		log.Printf("URL: %s\n", purl)
		res, err := client.Get(purl)
		if err != nil {
			log.Fatalf("Failed to make request(%q): %v", purl, err)
		}
		blob, err := ioutil.ReadAll(res.Body)
		log.Printf("Blob: %s\n", blob)
		if err != nil {
			log.Fatalf("%q failed to read body: %v", purl, err)
		}
		pauseTime := maxIdleTimeout / 4
		log.Printf("Pausing for %s\n\n", pauseTime)
		<-time.After(pauseTime)
	}
}
