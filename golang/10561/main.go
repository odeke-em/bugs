package main

type I interface {
	Foo()
}

type X int

func (*X) Foo() {}

func main() {
	var x X
	var i I = x
	switch i.(type) {
	case X:
	}
}
