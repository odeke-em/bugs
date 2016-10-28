package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	a := complex128(1)
	b := complex128(0)
	fmt.Printf("value=%v\n", cmplx.Cot(a/b))
}

// Cot => cos(X)/sin(X)
