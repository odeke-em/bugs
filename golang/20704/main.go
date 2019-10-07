package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"

	"golang.org/x/net/http2"
)

const (
	itemSize  = 1 << 10
	itemCount = 100
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < itemCount; i++ {
			_, err := w.Write(make([]byte, itemSize))
			if err != nil {
				return
			}
		}
	})

	srv := httptest.NewUnstartedServer(handler)
	srv.TLS = &tls.Config{
		NextProtos: []string{"h2"},
	}
	srv.StartTLS()

	cl := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	for i := 0; i < 10000; i++ {
		resp, err := cl.Get(srv.URL)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
}
