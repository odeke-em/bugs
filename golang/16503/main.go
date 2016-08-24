package main

import (
	"fmt"
)

var x int = y
var y int = z
var z int = y

func main() {
	fmt.Println(x, y)
}
