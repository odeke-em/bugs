package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	var l *int = nil

	rpc.Register(l)
	fmt.Println("Hello, playground")
}