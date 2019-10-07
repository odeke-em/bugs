package main

import (
	"net/url"
	"testing"
)

func copyValuesAfter(dst, src url.Values) {
	for k, vs := range src {
		dst[k] = append(dst[k], vs...)
	}
}

func copyValuesBefore(dst, src url.Values) {
	for k, vs := range src {
		for _, value := range vs {
			dst.Add(k, value)
		}
	}
}

func BenchmarkCopyValues_before(b *testing.B) {
	benchmarkCopyValues(b, copyValuesBefore)
}
func BenchmarkCopyValues_after(b *testing.B) {
	benchmarkCopyValues(b, copyValuesAfter)
}

var valuesCount int

func BenchmarkCopyValues(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dst := url.Values{"a": {"b"}, "b": {"2"}, "c": {"3"}, "d": {"4"}, "j": nil, "m": {"x"}}
		src := url.Values{
			"a": {"1", "2", "3", "4", "5"},
			"b": {"2", "2", "3", "4", "5"},
			"c": {"3", "2", "3", "4", "5"},
			"d": {"4", "2", "3", "4", "5"},
			"e": {"1", "1", "2", "3", "4", "5", "6", "7", "abcdef", "l", "a", "b", "c", "d", "z"},
			"j": {"1", "2"},
			"m": nil,
		}
		copyValues(dst, src)
		if valuesCount = len(dst["a"]); valuesCount != 6 {
			b.Fatalf(`%d items in dst["a"] but expected 6`, valuesCount)
		}
	}
	if valuesCount == 0 {
		b.Fatal("Benchmark wasn't run")
	}
}
