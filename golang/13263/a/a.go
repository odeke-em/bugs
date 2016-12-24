package a

var (
	a uint
	b = a
	c = uint64(b)
	d = uintptr(b)
)
