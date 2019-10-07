package main_test

import (
	"math/bits"
	"testing"
)

func TestDiv(t *testing.T) {
	v := int64(12345*1000000000+54321)
	div := int32(1000000000)

	hi, lo := uint32(v>>32), uint32(v&0xFFFFFFFF)
	quo, rem := bits.Div32(hi, lo, uint32(div))
	if g, w := quo, uint32(12345); g != w {
		t.Errorf("Quotient:\n\tgot  %d (0x%X)\n\twant %d (0x%X)", g, g, w, w)
	}
	if g, w := rem, uint32(54321); g != w {
		t.Errorf("Remainder :\n\tgot  %d (0x%X)\n\twant %d (0x%X)", g, g, w, w)
	}
}
