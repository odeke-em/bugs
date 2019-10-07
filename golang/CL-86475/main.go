package main

/*
#define foo 2
int fooz() {
  return foo;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Printf("%d\n", int(C.fooz()))
}
