package main

import (
    "./a"
    _ "./b"
)

type P string

func (p *P) Foo() error {
	panic("panic")
        return nil
}

func A() {
	var p P
        println(p.Foo)
}

func B() {
	var p P
	var f = p.Foo
	f()
}

func main() {
    if false {
        a.Bar()
    }
        var x P
        var f = x.Foo
        f()
}
