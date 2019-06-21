package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))
	defer cst.Close()

	client := cst.Client()
        tr := client.Transport.(*http.Transport)
        tr.IdleConnTimeout = 2500 * time.Millisecond
        log.Printf("Transport: %#v\n", tr)
        client.Transport = tr

	for {
		res, err := client.Get(cst.URL)
		if err != nil {
			log.Printf("Fetching request failed: %v", err)
			break
		}
		blob, _ := ioutil.ReadAll(res.Body)
		log.Printf("%s\n\n", blob)
		_ = res.Body.Close()
		<-time.After(5 * time.Second)
	}
}
