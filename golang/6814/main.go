package main

import "fmt"

type Foo struct {
	Bar string
}

type Lm interface {
	Baz() string
}

type Loo int

func PrintBar(lm Lm) {
	fmt.Println(Foo.Baz)
}

func main() {
	PrintBar(nil)
}
