package foo

import (
	"testing"
)

func TestFoo(t *testing.T) {
}

func (f *Foo) Name() string {
	return "foo:name"
}
