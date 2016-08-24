package main

import (
	"fmt"
	"runtime"
)

func f1() {
	var (
		w []*uint
		x = *w[0]
		y = x
		z = (uintptr)(y)
	)
	fmt.Printf("z: %v\n", z)
}

func f2() {
	var (
		x uint
		y = x
		z = uintptr(y)
	)
	fmt.Printf("z: %v\n", z)
}

func main() {
	fmt.Printf("Version: %s\n", runtime.Version())
}
