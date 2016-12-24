package main

func foo(args ...int) interface{} {
	return 0
}

func main() {
	var s1, s2, s3 string
	foo(
		1,
		s1,
		s2,
	)
}
