package main

import "fmt"

func main() {
	s := (-0 + 0i) + 0
	var zero complex128
	inf := 1.0 / zero
	fmt.Println(1/s, inf)
	if 1/s != inf {
		panic("1/negZeroCmplx != inf")
	}
}
