package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", time.Now())
	})
	addr := ":8888"
	println("Serving on " + addr)
	if err := http.ListenAndServeTLS(addr, "cert.pem", "key.pem", mux); err != nil {
		panic(err)
	}
}
