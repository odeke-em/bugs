package main

var A = func() int {
	return B()
}()

func B() int {
	return C()
}

func C() int {
	return D()
}

func D() int {
	if A == 0 {
		return 1
	}
	return 0
}

func main() {
}
