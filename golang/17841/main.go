package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/testOne", writeBackQuery)  // no slash at end or you lose the query string
	mux.HandleFunc("/testTwo/", writeBackQuery) // no slash at end or you lose the query string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()
	ts.Config.Handler = mux

	paths := [...]string{
		0: "/testOne?this=that",
		2: "/testTwo?foo=bar",
	}

	for i, path := range paths {
		res, err := http.Get(ts.URL + path)
		if err != nil {
			continue
		}
		slurp, _ := ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
		fmt.Printf("#%d: route: (%q) response: (%s)\n", i, path, slurp)
	}
}

func writeBackQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Query())
}
