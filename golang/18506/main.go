package main

import (
	"net/http"
	"log"
	"net/http/httputil"
)

func main(){

	req, err := http.NewRequest("GET", "https://start.exactonline.nl/api/v1/current/Me", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	rdata, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(rdata))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Status)
}
