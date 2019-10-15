package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to get an available socket: %v", err)
	}
	defer ln.Close()

	log.Printf("Server running at %s\n", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Failed to get a connection: %v", err)
		}
		go func(c net.Conn) {
			defer c.Close()
			serveContentBack(c)
		}(conn)
	}
}

func serveContentBack(w net.Conn) {
	buf := new(bytes.Buffer)
	gzw := gzip.NewWriter(buf)
	gzw.Write([]byte("<!doctype html><p>Hello</p>"))
        gzw.Close()
	fmt.Fprintf(w, "HTTP/1.1 200 OK\r\nContent-Encoding: gzip\r\nContent-Length: %x\r\n\r\n%s\r\n", buf.Len(), buf.String())
}
