/*
$ go version
go version go1.12.5 darwin/amd64

$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
BenchmarkIs408StringEqual-8       	2000000000	         1.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkIs408ByteEqualGlobal-8   	200000000	         8.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkIs408ByteEqualLocal-8         	100000000	        19.1 ns/op	       0 B/op	       0 allocs/op
*/

package main

import (
	"bytes"
	"testing"
)

func is408Message1(buf []byte) bool {
	if len(buf) < len("HTTP/1.x 408") {
		return false
	}
	if string(buf[:7]) != "HTTP/1." {
		return false
	}
	return string(buf[8:12]) == " 408"
}

func is408Message2(buf []byte) bool {
	if len(buf) < len("HTTP/1.x 408") {
		return false
	}
	if string(buf[:7]) != "HTTP/1." {
		return false
	}
	return string(buf[8:12]) == " 408"
}

var (
	b1 = []byte("HTTP/1.")
	b2 = []byte(" 408")
)

func is408Message3(buf []byte) bool {
	if len(buf) < len("HTTP/1.x 408") {
		return false
	}
	if !bytes.Equal(buf[:7], []byte("HTTP/1.")) {
		return false
	}
	return bytes.Equal(buf[8:12], []byte(" 408"))
}

func BenchmarkIs408StringEqual(b *testing.B) {
	buf := []byte("HTTP/1. foobar/wefwefojwalkjwelkjewf")
	var r bool
	for i := 0; i < b.N; i++ {
		r = is408Message1(buf)
	}
	sink = r
}
func BenchmarkIs408ByteEqualGlobal(b *testing.B) {
	buf := []byte("HTTP/1. foobar/wefwefojwalkjwelkjewf")
	var r bool
	for i := 0; i < b.N; i++ {
		r = is408Message2(buf)
	}
	sink = r
}
func BenchmarkIs408ByteEqualLocal(b *testing.B) {
	buf := []byte("HTTP/1. foobar/wefwefojwalkjwelkjewf")
	var r bool
	for i := 0; i < b.N; i++ {
		r = is408Message3(buf)
	}
	sink = r
}

var sink bool

func TestIs408(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"HTTP/1.0 408", true},
		{"HTTP/1.1 408", true},
		{"HTTP/1.8 408", true},
		{"HTTP/2.0 408", false}, // maybe h2c would do this? but false for now.
		{"HTTP/1.1 408 ", true},
		{"HTTP/1.1 40", false},
		{"http/1.0 408", false},
		{"HTTP/1-1 408", false},
	}
	for _, tt := range tests {
		if got := is408Message1([]byte(tt.in)); got != tt.want {
			t.Errorf("is408Message1(%q) = %v; want %v", tt.in, got, tt.want)
		}
		if got := is408Message2([]byte(tt.in)); got != tt.want {
			t.Errorf("is408Message2(%q) = %v; want %v", tt.in, got, tt.want)
		}
		if got := is408Message3([]byte(tt.in)); got != tt.want {
			t.Errorf("is408Message3(%q) = %v; want %v", tt.in, got, tt.want)
		}
	}
}
