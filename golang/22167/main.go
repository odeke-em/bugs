package p

var s string

func f(x, y, z int) int

func g() int {
	return f("hello", s, "world")
	return f(s, "hello", "world")
}
