package main

import (
	"log"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.example.com", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("1. %v", err)
	}

	err = res.Body.Close()
	if err != nil {
		log.Fatalf("2. %v", err)
	}

	// this is where things start to go sideways

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("3. %v", err)
	}

	log.Printf("Terminus!: %#v\n", res)
}
