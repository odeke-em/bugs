package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8833, "the port to run the server on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)

	http.HandleFunc("/a", handleA)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func handleA(rw http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		MaxAge: 3600,
		Name:   "a",
		Path:   "/a",
	}
	if c != nil {
	}
	rw.Header().Add("Set-Cookie", `name1="ab"`)
	rw.Header().Add("Set-Cookie", `name2="ab" foo`)
	rw.Header().Add("Set-Cookie", `name3=a b`)
	rw.Header().Add("Set-Cookie", `name4="3ab";am`)
	rw.Write([]byte("route a"))
}
