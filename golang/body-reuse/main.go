package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
)

func breadCrumbs(r *http.Request) (suffix string) {
	switch r.URL.Path {
	case "/b1":
		suffix = "/b2"
	case "/b2":
		suffix = "/b3"
	case "/b3":
		suffix = "/b4"
	case "/b4":
		suffix = "/b7"
	case "/b7":
		suffix = "/b8"
	}
	return suffix
}

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("path: %s\n", r.URL.Path)
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		w.Write(dump)
	}))
	defer ts.Close()

	var body string
	var method string
	flag.StringVar(&body, "body", "", "body to send")
	flag.StringVar(&method, "method", "POST", "method to use in client request")
	flag.Parse()

	req, err := http.NewRequest(method, ts.URL, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", slurp)
}
