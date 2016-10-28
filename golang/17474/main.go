package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var uri string
	flag.StringVar(&uri, "uri", "https://t.co/0SimgD6auP", "uri to follow location on")
	flag.Parse()

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("req: %v; via: %v\n", req, via)
			return fmt.Errorf("done")
		},
	}

	res, err := client.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
