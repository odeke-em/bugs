package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestIt(t *testing.T) {
	const fakeConnectionToken = "X-Fake-Connection-Token"
	const backendResponse = "I am the backend"

	// someConnHeader is some arbitrary header to be declared as a hop-by-hop header
	// in the Request's Connection header.
	const someConnHeader = "X-Some-Conn-Header"

	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c := r.Header.Get(fakeConnectionToken); c != "" {
			t.Errorf("handler got header %q = %q; want empty", fakeConnectionToken, c)
		}
		if c := r.Header.Get(someConnHeader); c != "" {
			t.Errorf("handler got header %q = %q; want empty", someConnHeader, c)
		}
		w.Header().Set("Connection", someConnHeader+", "+fakeConnectionToken)
		w.Header().Set(someConnHeader, "should be deleted")
		w.Header().Set(fakeConnectionToken, "should be deleted")
		io.WriteString(w, backendResponse)
	}))
	defer backend.Close()
	backendURL, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatal(err)
	}
	proxyHandler := httputil.NewSingleHostReverseProxy(backendURL)
	frontend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyHandler.ServeHTTP(w, r)
	}))
	defer frontend.Close()

	getReq, _ := http.NewRequest("GET", frontend.URL, nil)
	res, err := frontend.Client().Do(getReq)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body: %v", err)
	}
	if got, want := string(bodyBytes), backendResponse; got != want {
		t.Errorf("got body %q; want %q", got, want)
	}
	if c := res.Header.Get(someConnHeader); c != "" {
		t.Errorf("handler got header %q = %q; want empty", someConnHeader, c)
	}
	if c := res.Header.Get(fakeConnectionToken); c != "" {
		t.Errorf("handler got header %q = %q; want empty", fakeConnectionToken, c)
	}
        t.Logf("Res.Header: %v\n", res.Header)
}
