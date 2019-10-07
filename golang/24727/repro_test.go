package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"net"
	"testing"
)

func TestWriteCount(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to bind to an available address: %v", err)
	}
	defer ln.Close()

	type result struct {
		n   int64
		err error
	}

	resultsChan := make(chan *result, 1)
	go func() {
		defer close(resultsChan)

		rConn, err := ln.Accept()
		if err != nil {
			t.Fatalf("Failed to accept a connection: %v", err)
		}
		n, err := io.Copy(ioutil.Discard, rConn)
		resultsChan <- &result{n: n, err: err}
	}()

	addr := ln.Addr().String()
	wConn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("Failed to dial to address (%s): %v", addr, err)
	}
	wn, err := io.Copy(wConn, io.LimitReader(rand.Reader, 1<<20))
	if err != nil {
		t.Fatalf("Failed to write out bytes, error: %v", err)
	}
	_ = wConn.Close()
	res := <-resultsChan
	t.Logf("ServerReadCount %d ClientWriteCount %d", res.n, wn)
	if g, w := res.n, wn; g != w {
		t.Errorf("Byte counts: serverReadCount %d clientWriteCount %d", g, w)
	}
	if res.err != nil {
		t.Fatalf("Unexpected error from server: %v", err)
	}
}
