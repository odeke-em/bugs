package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	addr := ":9789"
	log.Printf("Serving at: %s", addr)
	http.HandleFunc("/", helloWorld)
	if err := http.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	defer r.Body.Close()
	log.Printf("%s\n\n", dump)
	w.Write(dump)
}
