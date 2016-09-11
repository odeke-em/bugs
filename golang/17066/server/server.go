package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func echoBack(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	slurp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("gotBody: %s from %v\n", slurp, req.RemoteAddr)

	rw.Write(slurp)
}

func main() {
	server := new(http.Server)
	server.Addr = ":5789"

	http.HandleFunc("/", echoBack)
	log.Printf("Serving on %s\n", server.Addr)

	if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}
}
