package main

import "fmt"

//go:noline
func foo() (bool, string) {
	return true, "Hello"
}

func main() {
	_, s := foo()
	key := []byte(s)
	key = key[:32]
	fmt.Println(len(key))
	fmt.Println(key)
}
