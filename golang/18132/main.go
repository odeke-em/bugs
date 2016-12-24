package main

import (
	"net/http"
	"log"
)

func main() {
	resp, _ := http.Get("https://github.com/golang/go/issues/18132")
	defer resp.Body.Close()
	log.Printf("headers; %#v\n", resp.Header)
	log.Println(resp.ContentLength)
	log.Println(resp.Uncompressed)
}
