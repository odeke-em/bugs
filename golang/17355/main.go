package main

import (
	"io"
	"log"
	"net/http/httputil"
)

func main() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("7\r\n12345678"))
	}()

	cr := httputil.NewChunkedReader(pr)
	readBuf := make([]byte, 7)
	n, err := cr.Read(readBuf)
	log.Printf("Read %v, %v: %q", n, err, readBuf[:n])
}
