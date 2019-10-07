package main

import (
	"fmt"
	"net/http"
)

var securityPreservingHTTPClient = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) > 0 && via[0].URL.Scheme == "https" && req.URL.Scheme != "https" {
			lastHop := via[len(via)-1].URL
			return fmt.Errorf("redirected from secure URL %s to insecure URL %s", lastHop, req.URL)
		}
		return nil
	},
}

func main() {
	_, err := securityPreservingHTTPClient.Get("https://vcs-test.golang.org/insecure/go/insecure")
	fmt.Println(err)
}
