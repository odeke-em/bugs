package main

import (
	"fmt"
)

type MyString string

func main() {
	a := MyString("sss")
	b := []rune(a) 
	fmt.Println(b)
}
