package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 256; i++ {
		go func(n int) {
			ch <- fmt.Sprintf("string %d", n)
		}(i)
	}
	time.Sleep(time.Second)
	close(ch)
}
