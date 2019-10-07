package main

import "testing"

type A struct{}

func (a *A) Foo() {
	panic("here")
}

func TestFoo(t *testing.T) {
	var a A
	println(a.Foo)
}

func TestBar(t *testing.T) {
	var a A
	var f = a.Foo
	f()
}