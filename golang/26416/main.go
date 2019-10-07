package main

type Y struct {
	X int
}

type Z struct {
	*Y
}

type A struct {
	*Z
}

type B struct {
	*A
}

func main() {
	_ = B{X: 0}
}
