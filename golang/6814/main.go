package main

import "fmt"

type Foo struct {
	Bar string
}

func PrintBar(f Foo) {
	fmt.Println(Foo.Bar)
}

func main() {
	f := Foo{"baz"}
	PrintBar(f)
}
