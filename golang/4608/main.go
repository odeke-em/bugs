package main

import (
	"fmt"
	"runtime"
)

type L struct {
	x, y int
}

var x L

func main() {
	println(runtime.Version())
	foo(0)
	fmt.Println(x)
}

func foo(bar int) {
	defer func() {
		_ = recover()
	}()
	x = L{1, 1 / bar}
}
