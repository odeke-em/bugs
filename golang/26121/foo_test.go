package wrongline

import (
	"testing"
)

type wrap struct {
	testing.TB // Also fails with *testing.T
}

func (w wrap) fail() {
	w.Helper()
	w.Error("fail")
}

type helper interface {
	testing.TB
	fail()
}

func withHelper(h helper) {
	h.Helper()
	h.fail() // Error is logged from here
}

func TestLineNumber(t *testing.T) {
	withHelper(wrap{t})
}