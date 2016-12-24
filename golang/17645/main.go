package main

type Foo struct {
	X int
}

func main() {
	var s []int
	var _ string = append(s, Foo{""})
	_ = string(Foo{""}, "hey")
}
