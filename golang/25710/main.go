package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/b", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("handle /b"))
	})
	mux.HandleFunc("/b/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("handle /b/"))
	})

	uri, _ := url.Parse("http://localhost:9000/b")
	mux.Handle("/api/a", http.StripPrefix("/api/a", httputil.NewSingleHostReverseProxy(uri)))

	cst := httptest.NewServer(mux)
	defer cst.Close()

	// Then for the client.
	res, err := http.Get(cst.URL + "/api/a")
	if err != nil {
		log.Fatalf("Failed to perform request: %v", err)
	}
	bgot, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	got := string(bgot)
	if want := "handle /b/"; got != want {
		log.Fatalf("Got: %q\nWant:%q\n", got, want)
	}
}
