package main

import "fmt"

type T struct{ V, W, X int }

func main() {
	var tp *int

	t := &T{
		V: *tp,
		X: 5,
		W: 0,
	}
	fmt.Println(t)
}
