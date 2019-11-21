package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("olleh"))
	})
	addr := ":8888"
	println("Serving on " + addr)
	if err := http.ListenAndServeTLS(addr, "cert.pem", "key.pem", mux); err != nil {
		panic(err)
	}
}
