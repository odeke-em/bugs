package main

/*
#include <unistd.h>
void foo(uint32_t i) {}
*/
import "C"

func main() {
	i := int(5)
	C.foo(i)
}
