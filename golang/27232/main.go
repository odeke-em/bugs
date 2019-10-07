package p

type F = func(T) // compiles if not type alias or moved below T interface

type T interface {
	m(F)
}

type t struct{}

func (t) m(F) {}

var _ T = &t{}
