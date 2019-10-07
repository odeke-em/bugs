package main

import (
	"fmt"
	"runtime"
)

type string = []int

func main() {
	fmt.Println(runtime.Version())
}
