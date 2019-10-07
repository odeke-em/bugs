package p

type S struct {
	i interface {
		f() [1 << 27]S
	}
}