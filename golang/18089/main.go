package main

type T struct {
	x int
	_ int
}

func main() {
	x := T{0, 0}
	_ = x
}
