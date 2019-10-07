package foo

import (
	"reflect"
	"testing"
)

var want [0]byte

func TestIssue32595(t *testing.T) {
	deref := reflect.ValueOf(new([0]byte)).Elem().Interface()
	if g, w := reflect.DeepEqual(deref, want), true; g != w {
		t.Fatalf("got=%t want=%t", g, w)
	}
}
