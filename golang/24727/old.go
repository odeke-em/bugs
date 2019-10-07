package main

import (
	"bytes"
	"io"
	"net"
	"testing"
	"time"
)

func TestIssue24727(t *testing.T) {
	ln, err := net.Listen("tcp4", ":0")
	if err != nil {
		t.Fatalf("Failed to bind to an available address: %v", err)
	}
	defer ln.Close()

	type result struct {
		n   int64
		err error
	}

	resultsChan := make(chan *result)
	go func() {
		defer close(resultsChan)

		var buf [0x100000]byte
		rConn, rErr := ln.Accept()
		if rErr != nil {
			t.Fatalf("Failed to accept a connection: %v", rErr)
		}

		var n int64
		var in int
		var err error
		defer func() {
			resultsChan <- &result{n: n, err: err}
		}()

		for {
			in, err = rConn.Read(buf[:])
			n += int64(in)
			if err == nil {
				continue
			}

			ne, ok := err.(net.Error)
			if ok && ne.Timeout() {
				err = nil
			} else {
				return
			}
		}
	}()

	addr := ln.Addr().String()
	wConn, err := net.Dial("tcp4", addr)
	if err != nil {
		t.Fatalf("Failed to dial to address (%s): %v", addr, err)
	}

	bufSize := 4096
	aSlice := bytes.Repeat([]byte("a"), bufSize)
	var wn int64

	runClientWrite := func() {
		defer wConn.Close()

		for i := 0; i < 1<<20; i += bufSize {
			wConn.SetWriteDeadline(time.Now().Add(1 * time.Millisecond))
			n, err := wConn.Write(aSlice)
			if err != nil {
				if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
					err = nil
					if n != 0 {
						t.Errorf("n %d on timeout err: %v", n, err)
					} else {
						t.Errorf("Error: %v", err)
						return
					}
				}
			}
			wn += int64(n)
		}
	}
	runClientWrite()

	res := <-resultsChan
	t.Logf("ServerReadCount %d ClientWriteCount %d", res.n, wn)
	if g, w := res.n, wn; g != w {
		t.Errorf("Byte counts: serverReadCount %d clientWriteCount %d", g, w)
	}
	if err := res.err; err != nil && err != io.EOF {
		t.Fatalf("Unexpected error from server: %+v res: %#v", err, res)
	}
}
