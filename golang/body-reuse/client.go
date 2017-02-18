package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	var payload string
	var method string
	flag.StringVar(&payload, "payload", "", "the payload of the request")
	flag.StringVar(&method, "method", "POST", "the method of the request")
	flag.Parse()

	req, err := http.NewRequest(method, "http://localhost:8080/", strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	slurp, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("slurp: %s\n", slurp)
}
