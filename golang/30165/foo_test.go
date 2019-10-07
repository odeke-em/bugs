package main

import (
	"strings"
	"testing"
)

var urls = []string{"", "   ", "/ ", "/x", " / ", "fooooooooobar"}

func BenchmarkHasSuffix(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, url := range urls {
			ok := false
			if !strings.HasSuffix(url, "/") {
				ok = true
			}
			if !ok {
				b.Fatalf("%q returned false", url)
			}
		}
	}
}

func BenchmarkPlain(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, url := range urls {
			ok := false
			if url == "" || url[len(url)-1] != '/' {
				ok = true
			}
			if !ok {
				b.Fatalf("%q returned false", url)
			}
		}
	}
}
