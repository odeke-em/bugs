package main

func foo() (string, interface{}) {
	return "%s", 10
}

func bar(s string, v interface{}) {
	_ = s
	_ = v

	var prod uint64
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			prod += uint64(i * j)
		}
	}

	if prod < 100 {
		panic("Failure")
	}
	if prod == 1e9 {
		println("Perfection")
	}
}

func main() {
	bar(foo())
}
