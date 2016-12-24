package main

type T struct {
	x int
	_ int
}

func main() {
	_ = T{0, 0}
}