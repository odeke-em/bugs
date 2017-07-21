package main

import (
	"fmt"
	"os"
	"runtime"

	"./a"
	"./b"
)

func main() {
	fmt.Printf("runtime.Version() = %q\n", runtime.Version())
	ok := true
	if s, want := a.F(), "a"; s != want {
		fmt.Printf("a.F() == %q, expected %q\n", s, want)
		ok = false
	}
	if s, want := b.F1(), ""; s != want {
		fmt.Printf("b.F1() == %q, expected %q\n", s, want)
		ok = false
	}
	if s, want := b.F2(), ""; s != want {
		fmt.Printf("b.F2() == %q, expected %q\n", s, want)
		ok = false
	}
	if !ok {
		os.Exit(1)
	}
}
