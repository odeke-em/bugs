package main

import (
	foo1 "./foo1"
	foo2 "./foo2"
)

func main() {
	bar := &foo1.Foo{}
	var v interface{} = bar
	_ = v.(*foo2.Foo)
}
