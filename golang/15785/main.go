package main

import (
	"fmt"
)

type test struct {
	a, b, c, d int
}

func params() (int, int, int, int) {
	return 1, 2, 3, 4
}

func f(a, b, c, d int) {
}

func main() {
	// Not Working
	a := []int{params()}
	b := test{params()}

	// Works fine
	f(params())

	fmt.Println(a, b)
}
