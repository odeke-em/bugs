package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, strings.NewReader("HTTP/1.1 400 Free\r\n"+
			"Content-Length: 123\r\n"+
			"\r\n"+
			"aloha ola bonjour"))
	}))
	defer cst.Close()

	res, err := cst.Client().Get(cst.URL)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	fmt.Printf("res: %+v\n", res)
}
