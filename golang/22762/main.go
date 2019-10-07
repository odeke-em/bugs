package main

import (
"log"
	"time"
)

var safeSlice = []int{1}
var testSlice = make([]int, 0, 1)

func main() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			for {
				s1 := append(safeSlice, index)
log.Printf("s1: %v\n", s1)
				if s1[1] != index {
					panic("")
				}
				s2 := append(testSlice, index) // <-   data race here
				if s2[0] != index {            // <-   panic here
					panic("")
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 10)
}
