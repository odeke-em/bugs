package main

import (
	"fmt"
)

func main() {
	s := "123"
	const t = "123"
	fmt.Println("" < s)
	fmt.Println("" < t)
}