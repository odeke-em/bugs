package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

var _2HopsMaxClient = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) > 2 {
			return fmt.Errorf("max 2 hops")
		}
		return nil
	},
}

func main() {
	cst1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}))
	defer cst1.Close()

	cst2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", cst1.URL)
		w.WriteHeader(308)
	}))
	defer cst2.Close()

	cst3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", cst2.URL)
		w.WriteHeader(308)
	}))
	defer cst3.Close()

	cst4 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", cst3.URL)
		w.WriteHeader(308)
	}))
	defer cst4.Close()
        println("original URL: ", cst4.URL)

	if _, err := _2HopsMaxClient.Get(cst4.URL); err != nil {
		log.Fatalf("Failed to fetch URL: %v\n", err)
	}
}
