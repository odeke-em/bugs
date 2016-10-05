package main

import (
	"net/http"
)

func echo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	w.Write([]byte("echo"))
}

func main() {
	http.HandleFunc("/", echo)

	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
