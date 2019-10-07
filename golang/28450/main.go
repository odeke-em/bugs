package main

import (
	"fmt"
)

func Foo(bar, foo, fizz, args ...interface{}) {
}

func main() {
	fmt.Println("Hello, playground")
	Foo("fee", 5 "fogh" `fum`)
}
