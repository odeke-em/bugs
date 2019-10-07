package main

import (
	"fmt"
)

type myStruct struct{}

func main() {
	if a := myStruct{}; false {
		fmt.Printf("oh %s", a)
	}
}
