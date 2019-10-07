package main

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var urllist = []string{
	// redacted
	"https://golang.org",
	"https://orijtech.com",
	"https://google.com",
	"https://twitter.com",
	"https://example.com",
	"https://apple.com",
	"https://gmail.com",
	"https://drive.google.com",
	"https://medium.com",
}

const nconcurrent = 5

func main() {
	t := http.Transport{
		TLSHandshakeTimeout:   20 * time.Second,
		ResponseHeaderTimeout: 1 * time.Minute,
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   20 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).Dial,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(nconcurrent)
	for i := 0; i < nconcurrent; i++ {
		go func() {
			for urlstr := range ch {
				req, err := http.NewRequest("GET", urlstr, nil)
				if err != nil {
					log.Fatal(err)
				}

				resp, err := t.RoundTrip(req)
				if err != nil {
					log.Println("RoundTrip", req.URL, err)
				} else {
					log.Println("RoundTrip", req.URL, resp.Status, resp.Proto)
					if resp.Body != nil {
						b := io.LimitReader(resp.Body, 10000000)
						_, err = ioutil.ReadAll(b)
						if err != nil {
							log.Println("ReadAll", req.URL, err)
						}
					}
				}
			}

			wg.Done()
		}()
	}

	for _, urlstr := range urllist {
		ch <- urlstr
	}

	log.Println("Feed completed")
	close(ch)
	wg.Wait()
}
