package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

const (
	header = "Fees"
	value  = "application/json"
	addr   = "localhost:8080"
)

func testHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(header, value)
	w.WriteHeader(http.StatusOK)
}

func testWrongOrderHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(header, value)
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/testWrongOrder", testWrongOrderHandler)
	go http.ListenAndServe(addr, nil)

	checkHeaderValue("/test")
	checkHeaderValue("/testWrongOrder")
}

func checkHeaderValue(path string) {
	req, _ := http.NewRequest("GET", path, nil)
	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)
	res := rec.Result()
	fmt.Println(path, "(httptest)", res.Header[header])

	resp, err := http.Get("http://" + addr + path)
	if err != nil {
		panic(err)
	}
	fmt.Println(path, "(http.Get)", resp.Header.Get(header))
}
