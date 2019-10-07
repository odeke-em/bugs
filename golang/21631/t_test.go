package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	f(t)
	g(t)
}

func f(tb testing.TB) {
	tb.Helper()
	tb.Error("asdf")
}

func g(t *testing.T) {
	t.Helper()
	t.Error("asdf")
}
