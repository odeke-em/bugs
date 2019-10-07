package main

import (
	"fmt"
)

func main() {
	b := make([]byte, 0)
	b = append(b, 1)
	fmt.Println(len(b))
	fmt.Println(len(b) - 2)
	s := string(b[1 : len(b)-2])
	fmt.Println(s)
}
