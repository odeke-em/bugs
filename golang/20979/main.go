package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://precious.jp")
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	if !statusOK(res.StatusCode) {
		log.Fatalf("res: %s", res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}

func statusOK(code int) bool { return code >= 200 && code <= 299 }
