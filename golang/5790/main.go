package main

import "fmt"

func main() {
	b := []int16{2, 0}
	_ = b[0] / b[1]
	var x int
	fmt.Printf("still alive: %v\n", x)
}
