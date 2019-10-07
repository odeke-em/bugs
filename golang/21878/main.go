package main

/*
static void return_void() { return; }
*/
import "C"

func main() {
	x := C.return_void() // ERROR HERE
	var y C.void = x     // ERROR HERE
	var z *C.void = &y   // ERROR HERE
	var b [0]byte = *z   // ERROR HERE
	_ = b
}
