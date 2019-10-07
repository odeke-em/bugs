package main

type Foo struct {
	A int
	B int
	C interface{}
	Bar
}

type Bar struct {
	A string
}

func main() {
	f := Foo{
		1,
		2,
		Bar{"hello", "A"},
	}
}
