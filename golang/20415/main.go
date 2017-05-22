package main

func f() {
	_ = g[:]
}

var g [1]byte

func f2() {
	for g := 0; g < 2; g++ {
		go func(g int) {
		}(g)
	}
}

var g interface{}

func main() {
}
