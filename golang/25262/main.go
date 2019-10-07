package main

func main() {}

type I interface {
	M1() interface {
		I
	}
	M2() interface {
		I
	}
}