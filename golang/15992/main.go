package main

import (
	"fmt"
)

func f(a []byte) ([]byte, []byte) {
	return a, []byte("abc")
}

func g(a []byte) ([]byte, string, int) {
	return a, "abc", 10
}

func main() {
	a := []byte{1, 2, 3}
	n := copy(f(a))
	fmt.Println(n, a)

	b := []byte{1, 2, 3}
	n = copy(g(b))
	fmt.Println(n, b)

	v := copy(func(a []byte) ([]byte, []byte) {
		return a, []byte("abc")
	}(a))
	fmt.Printf("v: %v\n", v)
}
