package main

import (
	"strings"
	"testing"
)

func BenchmarkFooNoSubTests(b *testing.B) {
	b.Logf("foo-outer")
	tests := []string{" aaaaa", "    aaaaa      ", "aaaaa   "}
        for _, t := range tests[:1] {
			got := strings.TrimSpace(t)
			if got != "aaaaa" {
				b.Fatalf("got back %q\n", got)
			}
	}
}
