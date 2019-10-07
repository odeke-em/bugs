package main

import (
	"fmt"
	"time"
)

func main() {
	writeTimer := time.NewTimer(10 * time.Millisecond)
	for i := 0; i < 100; i++ {
		fmt.Printf("#%d\n", i)
		select {
		case <-writeTimer.C:
		default:
		}
	}
}
