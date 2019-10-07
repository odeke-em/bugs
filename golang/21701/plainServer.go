package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8877, "the port to listen on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)
	log.Printf("server listening at %q\n", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("listen: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("listen err: %v", err)
			continue
		}
		go handleConn(conn)
	}
}

var malformed = []byte(
`HTTP/1.1 400
Date: Wed, 30 Aug 2017 19:09:27 GMT
Content-Type: text/html; charset=utf-8
Content-Length: 10
Connection: close
Last-Modified: Wed, 30 Aug 2017 19:02:02 GMT
Vary: Accept-Encoding
Server: cloudflare-nginx` + "\r\n\r\nAloha Olaa")

func handleConn(conn net.Conn) {
	conn.Write(malformed)
	conn.Close()
}
