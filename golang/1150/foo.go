package main

import (
	"fmt"
	"math"
)

func foo() {
	var f float64 = 18446744073709552000
	fmt.Printf("%.0f\n%d\n%d\n", f, uint64(f), uint64(math.MaxUint64))
}
