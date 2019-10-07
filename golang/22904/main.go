package main

import "fmt"

type a struct {
	b
}

type b struct {
	a
}

func main() {
	fmt.Println(a{})
}