package a1

type Closer interface {
	Close() error
}

func NewCloser() Closer {
	return nil
}
