package main

import (
	"fmt"
)

var x int = y

var y int = x

func main() {
	fmt.Println(x, y)
}
