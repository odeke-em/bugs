package main

import (
	"compress/gzip"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func compressSize(r io.Reader, level int) int64 {
	prc, pwc := io.Pipe()
	go func() {
		defer pwc.Close()
		gzw, err := gzip.NewWriterLevel(pwc, level)
		if err != nil {
			log.Printf("level: #%d err: %v", level, err)
		}
		io.Copy(gzw, r)
		gzw.Flush()
		gzw.Close()
	}()
	n, _ := io.Copy(ioutil.Discard, prc)
	return n
}

func main() {
	r := io.LimitReader(rand.Reader, 100000)
	for level := 1; level <= 9; level++ {
		size := compressSize(r, level)
		fmt.Printf("Level: %d size: %d\n", level, size)
	}
}
