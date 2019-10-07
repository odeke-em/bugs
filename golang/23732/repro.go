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
	_ = Foo{
		1,
		2,
		3,
	}

	_ = Foo{
		1,
		2,
		3,
		Bar{"A", "B"},
	}

	_ = Foo{
		1,
		2,
		Bar{"A", "B"},
	}
}
