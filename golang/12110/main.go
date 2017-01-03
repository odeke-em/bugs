package main

type S struct {
	boo T
}

type T struct {
	a int64
	b string
}

func main() {
	s := &S{
		baz: T{
			a: 1,
			b: "foo",
		},
	}
}
