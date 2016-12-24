package main

import (
	"fmt"
	"runtime"
)

var (
	a uint
	b = a
	c = (uint64)(b)
)

func main() {
	fmt.Printf("Version: %s\n", runtime.Version())
}
