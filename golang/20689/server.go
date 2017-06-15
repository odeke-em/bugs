package main

import (
	"log"
	"net/http"
	"strings"
)

func withLargeHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-wimmie", strings.Repeat("wimmie nah", 100001))
}

func main() {
	http.HandleFunc("/", withLargeHeader)

	err := http.ListenAndServeTLS(":9999", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
