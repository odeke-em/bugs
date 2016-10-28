package main

import "fmt"

func main() {
/*
if false {
	_, x := X()
	fmt.Printf("x = %v\n", x)
}
*/

	_, y := Y()
	fmt.Printf("y = %v\n", y)
}

/*
func X() (i int, ok bool) {
	ii := int(1)
	return ii, 0 <= ii && ii <= 0x7fffffff
}
*/

func Y() (int, bool) {
	ii := int(1)
	return ii, 0 <= ii && ii < 20
}
