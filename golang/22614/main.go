package main

import (
	"bytes"
	"crypto/rand"
	"os"
	"time"
)

func main() {
	pr, pw, _ := os.Pipe()

	buf := new(bytes.Buffer)
	ostdin = os.Stdin
	os.Stdin = pr
	var orr io.Reader = rand.Reader
	defer func() {
		os.Stdin = ostdin
		rand.Reader = orr
	}()

        tr := &timedReader{File: pr, wait: 61 * time.Second}
        buf := make([]byte, 32)
}

type timedReader struct {
	*File
	wait time.Duration
}

func (tr timedReader) Read(b []byte) (int, error) {
	<-time.After(time.Duration(tr.wait))
	return len(b), nil
}
