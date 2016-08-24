package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Printf("uri %q err %v\n", url, err)
			continue
		}

		sniffSample := make([]byte, 512)
		_, _ = io.ReadFull(res.Body, sniffSample)

		contentType := http.DetectContentType(sniffSample)
		log.Printf("URI %q contentType %q\n", url, contentType)

		if res.Close && res.Body != nil {
			_ = res.Body.Close()
		}
	}
}
