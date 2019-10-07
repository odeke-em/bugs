package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 1 * time.Second,
			DualStack: true,
		}).DialContext,
	}}
	i := 0
	for {
		i++
		log.Printf("Making request #%d", i)
		res, err := client.Get("https://golang.org/pkg/net/http/#Transport")
		if err != nil {
			log.Fatalf("Error: %T error: %v\n", err, err)
		}
		io.Copy(ioutil.Discard, res.Body)
		_ = res.Body.Close()
		log.Println("Going to sleep")
		<-time.After(5 * time.Second)
	}
}
