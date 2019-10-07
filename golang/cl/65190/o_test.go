package main_test

import (
"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckContentTypeOnHead(t *testing.T) {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("r") == "" {
fmt.Printf("q: %v\n", q)
			w.Header().Set("Location", "/foo?r=1")
			w.WriteHeader(301)
			return
		}
fmt.Printf("done here\n")
	}))
	defer cst.Close()

	client := cst.Client()
	res, err := client.Head(cst.URL)
	if err != nil {
		t.Fatal(err)
	}
	if g, w := res.Header.Get("Content-Type"), "text/html; charset=utf-8"; g != w {
		t.Errorf("Content-Type: got=%q want=%q", g, w)
	}
	if cl := res.ContentLength; cl > 0 {
		t.Errorf("ContentLength: got=%d want <= 0", cl)
	}
}
