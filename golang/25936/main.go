package main

import (
	"fmt"
)

const (
	key1 = ""
	key2 = "x"
)

func main() {
	m := make(map[string]int)
	m[key1] = 99
	delete(m, key1)
	n1 := m[key2]
	m[key2]++
	n2 := m[key2]
	if n2 != n1+1 {
		fmt.Printf("BUG!!!!! %v; incremented %v to %v\n", key2, n1, n2)
	}
}
