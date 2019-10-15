package main

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, err := url.Parse("http://bradleyjkemp:password@example.com")
	if err != nil {
		panic(err)
	}
	blob, _ := json.MarshalIndent(u, "", "  ")
	println(string(blob))
	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
	}
	blob, _ = httputil.DumpRequest(req, true)
	println(string(blob))
        res, _ := http.DefaultClient.Do(req)
	blob, _ = httputil.DumpResponse(res, false)
	println(string(blob))
}
