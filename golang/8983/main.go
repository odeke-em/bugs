package main

import (
	"./a1"
	"./a2"
)

type Closer int

func newOne1() Closer {
	return a1.NewCloser()
}

func newOne2() Closer {
	return a2.NilCloser()
}

func main() {
}
