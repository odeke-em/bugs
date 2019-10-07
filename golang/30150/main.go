package main

import "runtime"

type P string

func (p P) String() string {
	runtime.GC()
	runtime.GC()
	runtime.GC()
	zzz := "ZZZ"
	sink = zzz
	return string(p)
}

var sink interface{}

func main() {
	defer func() {
		panic(P("YYY"))
	}()
	panic(P("XXX"))
}