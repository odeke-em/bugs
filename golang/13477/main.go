package main

import (
	"fmt"
)

type A struct {
	Gist string
	Id   string
	B
}

type foo struct {
	Foo    int
	Foobar int
	*A
	l
}

type l struct{}

type B struct {
	Influence string
}

func (f *foo) F1() {
}

func main() {
	f := &foo{Foo: 1, Foobar: 2}
	fmt.Printf("gist: %s\n", f.gust)
	f.F2()
	f.Influenc
	f.True
}
