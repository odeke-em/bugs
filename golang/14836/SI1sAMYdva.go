package main

import (
	"log"
	"net/url"
	"os"
)

func main() {
	values := []string{
		"https://golang.org:", "https://golang.org:80a",
	}

	for _, val := range values {
		ux, err := url.Parse(val)
		if err != nil {
			log.Printf("err: %v", err)
			os.Exit(1)
		}
		log.Printf("resp: %v", ux.Host)
	}
}
