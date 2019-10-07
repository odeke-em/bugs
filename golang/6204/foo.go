package foo

type Foo struct {
}

func (f *Foo) String() string {
	return "foo"
}

func (f *Foo) All() string {
	return f.String() + f.Name()
}
