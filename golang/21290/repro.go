package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	log.SetFlags(0)
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, bufw, _ := w.(http.Hijacker).Hijack()
		conn.Write([]byte(
`HTTP/1.0 200 OK
Content-Length: 13
Content-Disposition = attachment; filename=file.asc


Hello, World!`))
                bufw.Flush()
                conn.Close()
	}))
	defer cst.Close()

	resp, err := http.Get(cst.URL)
	if err != nil {
		log.Fatalf("Failed to perform plain get: %v", err)
	}
	_, _ = ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
}
