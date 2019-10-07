package main

import (
	"fmt"
)

func main() {
	v := []int{1, 2, 3}
	fmt.Printf("vPtr: %p\n", v)
	v = append(v)
	fmt.Printf("vPtr: %p\n", v)
v = append(v, 2, 4)
	fmt.Printf("vPtr: %p\n", v)
}
