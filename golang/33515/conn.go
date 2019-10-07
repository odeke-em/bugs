package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"
)

func main() {
	baseDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 1 * time.Second,
		DualStack: true,
	}
	conn, err := tls.DialWithDialer(baseDialer, "tcp", "golang.org:443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatalf("Failed to dial to connection: %v", err)
	}
	defer conn.Close()
	log.Printf("RemoteAddr: %s\n", conn.RemoteAddr())

	var allGoroutinesWg sync.WaitGroup
	defer allGoroutinesWg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	allGoroutinesWg.Add(1)
	// Requests goroutine.
	go func() {
		defer allGoroutinesWg.Done()

		i := 0
		paths := []string{"/doc/tos.html", "/pkg/net/http%23Transport/", "/nonexistent"}
		for {
			for _, path := range paths {
				asStr := fmt.Sprintf("GET %s HTTP/1.1\r\nConnection: keep-alive\r\nHost: golang.org\r\n\r\n", path)
				n, err := conn.Write([]byte(asStr))
				log.Printf("\n\n\033[35mWrote %d bytes, error: %v\033[00m\n\n", n, err)

				select {
				case <-time.After(2800 * time.Millisecond):
					i++
					fmt.Printf("Making the next request: #%d", i)
				case <-ctx.Done():
					fmt.Printf("Exiting writing goroutine: %v", ctx.Err())
					return
				}
			}
		}
	}()

	// Responses are read here.
	br := bufio.NewReader(conn)
	for {
		res, err := http.ReadResponse(br, nil)
		if err != nil {
			log.Printf("\033[31mError: %v\033[00m", err)
			return
		}
		blob, _ := httputil.DumpResponse(res, true)
		fmt.Printf("\n\n\n%s\n\n\n", blob)
	}
}
