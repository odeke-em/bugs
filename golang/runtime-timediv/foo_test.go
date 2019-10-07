package main

import (
	"math/bits"
	"testing"
)

func BenchmarkAsPlus(b *testing.B) {
	benchmarkThem(b, withPlus)
}

func BenchmarkAsOr(b *testing.B) {
	benchmarkThem(b, withOr)
}

var sink int32
var sinku32 uint32

func BenchmarkAsDiv(b *testing.B) {
	v := int64(12345*1000000000 + 54321)
	div := int32(1000000000)

	var quo, rem uint32

	for i := 0; i < b.N; i++ {
		hi, lo := uint32(v>>32), uint32(v&0xFFFFFFFF)
		quo, rem = 0, 0
		quo, rem = bits.Div32(hi, lo, uint32(div))
	}

	if g, w := quo, uint32(12345); g != w {
		b.Fatalf("Got %d\nWant %d", g, w)
	}
	if g, w := rem, uint32(54321); g != w {
		b.Fatalf("Got %d\nWant %d", g, w)
	}

}

const (
	v   = int64(12345*1000000000 + 54321)
	div = int32(1000000000)
)

func benchmarkThem(b *testing.B, fn func(int64, int32) int32) {
	for i := 0; i < b.N; i++ {
		sink = fn(v, div)
	}
	if sink == 0 {
		b.Fatal("Benchmarks weren't run")
	}
}

func withPlus(v int64, div int32) int32 {
	res := int32(0)
	for bit := 30; bit >= 0; bit-- {
		if v >= int64(div)<<uint(bit) {
			v = v - (int64(div) << uint(bit))
			res += 1 << uint(bit)
		}
	}
	return res
}

func withOr(v int64, div int32) int32 {
	res := int32(0)
	for bit := 30; bit >= 0; bit-- {
		if v >= int64(div)<<uint(bit) {
			v = v - (int64(div) << uint(bit))
			res |= 1 << uint(bit)
		}
	}
	return res
}
