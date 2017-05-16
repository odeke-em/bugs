package main

import (
	"runtime/debug"
)

func recurse() {
	debug.SetMaxStack(10000)
	f1()
}

func f1() {
	f2()
}

func f2() {
	f3()
}

func f3() {
	f4()
}

func f4() {
	f1()
}

func main() {
	recurse()
}
