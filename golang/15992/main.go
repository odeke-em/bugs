package main

import (
	"fmt"
)

type s string

func (_ s) f(a []byte) ([]byte, []byte) {
	return a, []byte("abc")
}

func g(a []byte, n int) (int, string, int) {
	return len(a), "abc", 10
}

func p() ([]byte, []byte) {
    return []byte("a"), []byte("bc")
}

func main() {
	a := []byte{1, 2, 3}
	n := copy((s)("a").f(nil))
	fmt.Println(n, a)

	b := []byte{1, 2, 3}
	n = copy(g(b, 10))
	fmt.Println(n, b)

	v := copy(func(a []byte) ([]byte, []byte) {
		return a, []byte("abc")
	}(a))
	fmt.Printf("v: %v\n", v)

        copy(p())
}
