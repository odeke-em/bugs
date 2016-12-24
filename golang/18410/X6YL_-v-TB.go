package main

import (
	"fmt"
	"runtime"
)
type X struct {
	A, B []byte
}
func (x X) Print() {
	fmt.Printf("%p %p\n", x.A, x.B)
}
func caller(f func()) {
	f()
}
func main() {
	caller(X{A: []byte{}}.Print)
	fmt.Printf("runtime.Version: %s\n", runtime.Version())
}
