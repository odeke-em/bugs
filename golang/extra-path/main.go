package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

type rproxy struct {
	handler http.HandlerFunc
}

func (rp *rproxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp.handler(w, r)
}

func main() {
	otherBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", r.URL)
	}))
	defer otherBackend.Close()

	ts := httptest.NewServer(&rproxy{handler: func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, otherBackend.URL+"?foo=zar", http.StatusPermanentRedirect)
	}})
	defer ts.Close()

	res, err := ts.Client().Get(ts.URL + "?foo=bar")
	if err != nil {
		log.Fatal(err)
	}
	slurp, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("slurp: %s\n", slurp)

}
