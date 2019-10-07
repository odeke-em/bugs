package main

import "fmt"

func main() {
	asd := ^uint(0)
	fmt.Printf("typ: %T\n", asd)
	println(asd)

	// println(^uint(0))
}
