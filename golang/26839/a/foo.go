package a

type P string

func (p P) Foo() string {
	panic("done")
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

func Bar() {
	B()
}
