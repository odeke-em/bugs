package main

import "fmt"

//go:noinline
func f(x complex64) complex64 {
	return x + 0
}

func main() {
	var zero complex64
	var inf = 1.0 / zero
	var negZero = -1.0 / inf
        var ts = 1/negZero
        fmt.Println(1/f(negZero), inf, negZero, ts, 1/ts)
}
