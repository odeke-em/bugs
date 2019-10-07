package main

/*
int magic() {
    return 1049;
}
*/
import "C"

import "fmt"

func main() {
	var magic = C.magic()
	fmt.Printf("magic: %v\n", magic)
}
