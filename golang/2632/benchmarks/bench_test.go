package benchmarks

import (
	"bytes"
	"testing"
)

func BenchmarkRangeStringLoop(b *testing.B) {
	s := string(make([]byte, 4096))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {
			_ = s[j]
		}
	}
}

func BenchmarkRangeStringCast(b *testing.B) {
	s := string(make([]byte, 4096))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, _ = range []byte(s) {
		}
	}
}

func BenchmarkRangeBytes(b *testing.B) {
	bs := make([]byte, 4096)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, _ = range bs {
		}
	}
}

func BenchmarkStringConversions(b *testing.B) {
	bs := bytes.Repeat([]byte("aa"), 1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(string(bs)) > 0 {
		}
	}
}
