package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	go func() {
		_, err := w.Write([]byte("123"))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 4)
		_, err = w.Write([]byte("456"))
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		l, err := net.Listen("tcp", "localhost:1234")
		if err != nil {
			panic(err)
		}
		defer l.Close()
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		_, err = io.CopyN(conn, r, 3)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		time.Sleep(time.Second / 5)
		c, err := net.Dial("tcp", "localhost:1234")
		if err != nil {
			panic(err)
		}
		defer c.Close()
		_, err = io.Copy(os.Stdout, c)
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}()
	time.Sleep(time.Second / 2)

	deadline := time.Now().Add(time.Second)
	err = r.SetDeadline(deadline)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 4)
	n, err := r.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("got %v, deadline: %v, now: %v\n", string(buffer[:n]), deadline.Unix(), time.Now().Unix())
}