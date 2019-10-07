package main

type a int

func (_ a) do() {
	panic("panic")
}

func foo() {
	var x a
	_ = x.do
}

func bar() {
	var m a
	f := m.do
	f()
}

func main() {
	bar()
}
