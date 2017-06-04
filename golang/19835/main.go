package main

/*
#include <stdio.h>

static void invoke(void (*f)()) {
	f();
}

void print_hello() {
	printf("Hello, !");
}
*/
import "C"

func main() {
	C.invoke(C.print_hello)
}