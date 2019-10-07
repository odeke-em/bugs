package a

type P string

func (p *P) Foo() string {
	panic("done")
}
