package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkFooSubTests(b *testing.B) {
	b.Logf("foo-outer")
	tests := []string{" aaaaa", "    aaaaa      ", "aaaaa   "}
	for i := 0; i < b.N; i++ {
		for _, t := range tests {
			got := strings.TrimSpace(t)
			if got != "aaaaa" {
				b.Fatalf("got back %q\n", got)
			}
		}
	}

	for i, t := range tests[:1] {
		b.Run(fmt.Sprintf("%s-%d", t, i), func(bb *testing.B) {
			got := strings.TrimSpace(t)
			bb.Logf("foo-inner")
			if got != "aaaaa" {
				bb.Fatalf("got back %q\n", got)
			}
		})
	}
}
