package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU(), runtime.GOMAXPROCS(0))

	var a uint64 = 0
	go func() {
		for {
			atomic.AddUint64(&a, uint64(1))
			// fmt.Printf("adding a\n")
			time.Sleep(time.Second)
		}
	}()

	for {
		val := atomic.LoadUint64(&a)
		fmt.Println(val)
		time.Sleep(time.Second)
	}
}
