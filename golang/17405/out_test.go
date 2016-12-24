package main_test

import (
	"testing"
)

func BenchmarkTypeAssertSingleReturnValue(b *testing.B) {
	n := 42
	v := interface{}(n)
	for i := 0; i < b.N; i++ {
		a := v.(int)
		if a != n {
			b.Fatalf("unexpected a: %d. Expecting %d", a, n)
		}
	}
}

func BenchmarkTypeAssertTwoReturnValues(b *testing.B) {
	n := 42
	v := interface{}(n)
	for i := 0; i < b.N; i++ {
		a, ok := v.(int)
		if !ok {
			b.Fatalf("type assertion failed for v=%v", v)
		}
		if a != n {
			b.Fatalf("unexpected a: %d. Expecting %d", a, n)
		}
	}
}
