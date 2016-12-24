package main

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	cst := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		<-time.Tick(20 * time.Second)
	}))

	defer cst.Close()

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 1,
		},
	}
	done := make(chan bool)
n := 100
	for i := 0; i < n; i++ {
		go func() {
			rs := rand.Intn(20)
			<-time.Tick(time.Duration(rs) * time.Second)
			res, _ := client.Post(cst.URL, "text/plain", nil)
			_, _ = io.Copy(ioutil.Discard, res.Body)
			res.Body.Close()
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
