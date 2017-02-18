package main

import (
	"net/http"
	"net/http/httputil"
)

func receive(rw http.ResponseWriter, req *http.Request) {
	slurp, _ := httputil.DumpRequest(req, true)
	_ = req.Body.Close()
	rw.Write(slurp)
}

func main() {
	http.HandleFunc("/", receive)
	http.ListenAndServe(":8080", nil)
}
