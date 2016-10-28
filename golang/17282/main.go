package main

type foo interface {
	foo()
	bar()
	baz()
}

type test struct {}


func main() {
	var _ foo = test{}
}
