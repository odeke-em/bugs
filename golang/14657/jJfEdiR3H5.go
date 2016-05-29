package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	path := "Order.thrift"

	b := filepath.Base(path)
	fmt.Println(b) // Order.thrift

	e := filepath.Ext(path)
	fmt.Println(e) // .thrift

	n := strings.TrimRight(b, e)
	fmt.Println(n)
}
