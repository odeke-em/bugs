package a2

type Closer interface {
	Close() error
}

func NilCloser() Closer {
	return nil
}
