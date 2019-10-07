package main

import (
	"fmt"
)

const arraySize = 1

// comment above and uncomment below, will work as expected.
// const arraySize = 2

type foo struct {
	bar [arraySize]*int
}

func main() {
	ch := make(chan foo, 2)
	var a int
	var b [arraySize]*int
	b[0] = &a
	ch <- foo{bar: b}
	close(ch)

	for v := range ch {
		// uncomment below, will work as expected.
		// fmt.Println(v.bar[0])
		for i := 0; i < 1; i++ {
			fmt.Println(v.bar[0], b[0])
		}
	}
}