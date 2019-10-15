package main

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func makeTestServer() (addr string) {
	//get a random port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	addr = listener.Addr().String()

	listener.Close()

	go http.ListenAndServe(listener.Addr().String(), http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
		io.Copy(ioutil.Discard, request.Body)
		request.Body.Close()

	}))

	time.Sleep(time.Millisecond * 200)
	return "http://" + addr
}

func BenchmarkDoNotReuseRequest(b *testing.B) {
	addr := makeTestServer()
	dialer := net.Dialer{
		Timeout: 4 * time.Second,
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodPost, addr, strings.NewReader("hello"))
		if err != nil {
			b.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			b.Fatal(err)
		}
		if resp.Body != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
	}
}

func BenchmarkReuseRequest(b *testing.B) {
	addr := makeTestServer()
	dialer := net.Dialer{
		Timeout: 4 * time.Second,
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}

	u, _ := url.Parse(addr)
	b.ResetTimer()
	b.ReportAllocs()
	req, err := http.NewRequest(http.MethodPost, addr, nil)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		req.URL = u
                req.Body = ioutil.NopCloser(strings.NewReader("hello"))
		resp, err := client.Do(req)
		if err != nil {
			b.Fatal(err)
		}
		if resp.Body != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
	}
}
