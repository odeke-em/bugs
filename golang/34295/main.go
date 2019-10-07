package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBlob, _ := httputil.DumpRequest(r, true)
		w.Write(reqBlob)
	}))
	defer cst.Close()

	reqBody := bytes.NewBufferString(`{}`)
	req, err := http.NewRequest("POST", cst.URL, reqBody)
	if err != nil {
		log.Fatalf("Cannot create request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Cannot do: %v", err)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Cannot read body: %v", err)
	}
	log.Printf("Response Body:\n%s", resBody)
}
