package main

import (
	"fmt"
	"runtime"
)

var f func((i), _ int)

func main() {
	fmt.Printf("version: %s\n", runtime.Version())
}
