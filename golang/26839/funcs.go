package main

type P string

var _ = (*P).Foo

func (p *P) Foo() {
	panic("panic")
}

func SMD() {
    panic("SMD")
}

func A() {
	_ = SMD
        println(SMD)
}

func B() {
	var f = SMD
	f()
}

func main() {
        B()
}
