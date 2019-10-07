package main

type Interface interface {
	Method()
}

func UseInterface(i Interface) {}

type Struct struct{}

func main() {
	UseInterface({})
}
