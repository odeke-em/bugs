package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func noBodyReq(w http.ResponseWriter, r *http.Request) {
dump, _ := httputil.DumpRequest(r, true)
log.Printf("%s\n", dump)
	if g, min := r.ContentLength, int64(0); g > min {
		http.Error(w, fmt.Sprintf("got %d min %d", g, min), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Eureka!")
}

func main() {
	http.HandleFunc("/", noBodyReq)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
