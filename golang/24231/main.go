package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"golang.org/x/net/trace"
)

func main() {
	go func() {
		for {
			t := trace.New("family", "title")
			time.Sleep(200 * time.Millisecond)
			t.Finish()
		}
	}()

	// EITHER:
	//
	// Start the webserver and trigger by viewing "Active" requests.
	//
	// http.ListenAndServe("localhost:8888", nil)

	// OR:
	// Run the test server and trigger the race in code:
	m := http.NewServeMux()
	m.HandleFunc("/debug/requests", trace.Traces)

	s := httptest.NewServer(m)
	s.Client().Get(fmt.Sprintf("%v/debug/requests?fam=family&b=-1", s.URL))
}