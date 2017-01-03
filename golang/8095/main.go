package main

import "fmt"

type BigInterface interface {
	Foo()
	Bar()
	Baz()
}

type Type struct{}

func main() {
	a := Type{}
	b := BigInterface(a)
	fmt.Println(b)
}
