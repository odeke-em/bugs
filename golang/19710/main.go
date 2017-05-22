package main

import "fmt"

func main() {
	fmt.Println("should be 0:", len(foobar()))
}

func foobar() map[int]bool {
	m := map[int]bool{}
	for i := 0; i < 3; i++ {
		m[i] = true
		defer delete(m, i)
	}
	return m
}

func myDelete(m map[int]bool, k int) {
	delete(m, k)
}
