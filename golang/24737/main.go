package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// Set up a bad proxy.
	os.Setenv("HTTP_PROXY", "http://unknown.proxy.example.com/")

	// Run a web server on the unspecified interface.
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	listener, err := net.Listen("tcp4", "")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(listener.Addr())
	go http.Serve(listener, nil)
	time.Sleep(100 * time.Millisecond)

	// Try to talk to the web server.
	response, err := http.Get(fmt.Sprintf("http://%s/bar", listener.Addr().String()))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	log.Print(response.Status)
}
