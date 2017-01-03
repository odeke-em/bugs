package main

import (
	"C"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	C.function_that_does_not_exist()
}
