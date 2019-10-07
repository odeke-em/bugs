package main

import (
	"context"
	"log"
	"net"

	"golang.org/x/net/http2"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		log.Fatal(err)
	}
	t := &http2.Transport{}
	cconn, err := t.NewClientConn(conn)
	if err != nil {
		log.Fatal("newClientConn: %v", err)
	}
	if err := cconn.Ping(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
