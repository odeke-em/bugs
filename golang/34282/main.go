package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/http/httptest"
	"sync"
)

func gotConn(info httptrace.GotConnInfo) {
	_ = info.Conn.RemoteAddr
}

func doGet(uri string) {
	ctx := httptrace.WithClientTrace(context.Background(), &httptrace.ClientTrace{
		GotConn: gotConn,
	})

	req, _ := http.NewRequest("GET", uri, nil)
	req = req.WithContext(ctx)

	if _, err := http.DefaultClient.Do(req); err != nil {
		log.Println(err)
	}
}

func main() {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	defer cst.Close()

	var wg sync.WaitGroup
	n := 200
	defer wg.Wait()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doGet(cst.URL)
		}()
	}
}
