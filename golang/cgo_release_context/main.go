package main

/*
int foo() {
    return 10;
}
*/
import "C"

import "fmt"

func main() {
	fmt.Printf("%d\n", C.foo())
}
