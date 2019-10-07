package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("#%d\n", i)
		<-time.After(10 * time.Millisecond)
	}
}
