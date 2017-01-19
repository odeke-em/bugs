package main

type T [16]int

func f1() (t T) {
	return g()
}

func f2() T {
	return g()
}

func g() T {
	return T{1, 2, 3, 5}
}

func main() {
}
