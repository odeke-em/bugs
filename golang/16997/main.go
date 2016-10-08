package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

func main() {
	proxy := "http://2.2.2.2:80"
	tr := &http.Transport{
		Proxy: func(*http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		},
		Dial: (&net.Dialer{Timeout: 500 * time.Millisecond}).Dial,
	}
	c := http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", "http://golang.org", nil)
	_, err := c.Do(req)
	fmt.Printf("err: %v\n", err)
	uerr, ok := err.(*url.Error)
	if ok {
		nerr, ok := uerr.Err.(*net.OpError)
		fmt.Println(ok, nerr)
	}
}
