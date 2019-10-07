package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a int
	f := func() {
		start := time.Now()
		stack := make([]byte, 2048)
		_ = runtime.Stack(stack, true)
		fmt.Printf("%s\n", stack)
		time.Sleep(time.Second)
		for time.Since(start) < 3*time.Second {
			a = 5
		}
	}
	go f()
	f()
	println(a)
}
