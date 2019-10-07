package main

import (
	"fmt"
)

type Foo struct{}

func (f *Foo) Call(cb func(*Foo)) {
	cb(f)
}

func main() {
	f := &Foo{}

	f.Call(func(f) { // Type for f missing
		fmt.Println(f)
	})
}
