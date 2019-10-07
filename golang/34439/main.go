package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

func main() {
	timeout := 1 * time.Second // reached
	// timeout := 3 * time.Second // not reached

	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot) // STATUS -> 418
		io.WriteString(w, "line1\n")
		w.(http.Flusher).Flush()
		time.Sleep(2 * time.Second)
		io.WriteString(w, "line2\n")
	}
	ts := httptest.NewServer(http.TimeoutHandler(h, timeout, "TIMED OUT\n"))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response status:", resp.Status)
	_, err = io.Copy(os.Stdout, resp.Body)
	log.Print(err)
}
