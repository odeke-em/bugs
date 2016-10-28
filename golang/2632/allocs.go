package main

import (
	"bytes"
	"fmt"
	"math"
	"runtime"
	"strconv"
)

var (
	infShortBytes      = []byte("inf")
	infLongBytes       = []byte("infinity")
	infPlusShortBytes  = []byte("+inf")
	infMinusShortBytes = []byte("-inf")
	infMinusLongBytes  = []byte("-infinity")
	infPlusLongBytes   = []byte("+infinity")
	nanBytes           = []byte("nan")
)

func main() {
	s := "123.45"
	bs := []byte(s)
	r := bytes.NewBufferString(s)

	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	println(memStats.Alloc, memStats.Mallocs, "initial")

	_, _ = strconv.ParseFloat(s, 64)
	runtime.ReadMemStats(memStats)
	println(memStats.Alloc, memStats.Mallocs, "parseFloat")

	if false {
		var v float64
		fmt.Fscanf(r, "%f", &v)
		runtime.ReadMemStats(memStats)
		println(memStats.Alloc, "done with plain parse")
	}

	_, _ = strconv.ParseFloat(string(bs), 64)
	runtime.ReadMemStats(memStats)
	println(memStats.Alloc, memStats.Mallocs, "parseFloat64Bytes1")

	_, _ = strconv.ParseFloat(string(bs), 64)
	runtime.ReadMemStats(memStats)
	println(memStats.Alloc, memStats.Mallocs, "parseFloat64Bytes2")

	_, _ = strconv.ParseFloat(s, 64)
	runtime.ReadMemStats(memStats)
	println(memStats.Alloc, memStats.Mallocs, "parseFloat64")

	_ = equalIgnoreCaseBytes([]byte("abcd"), []byte("gfeh"))
	runtime.ReadMemStats(memStats)
	println(memStats.Alloc, memStats.Mallocs, "ignoredCases")

	_, _ = specialBytes([]byte("truuuu"))
	runtime.ReadMemStats(memStats)
	println(memStats.Alloc, memStats.Mallocs, "trruuu")
}

var caseBit = byte('a' - 'A')

func equalIgnoreCaseBytes(ba1, ba2 []byte) bool {
	if len(ba1) != len(ba2) {
		return false
	}
	for i := 0; i < len(ba1); i++ {
		c1 := ba1[i]
		if 'a' <= c1 && c1 <= 'z' {
			c1 ^= caseBit
		}
		c2 := ba2[i]
		if 'a' <= c2 && c2 <= 'z' {
			c2 ^= caseBit
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}

func specialBytes(ba []byte) (f float64, ok bool) {
	if len(ba) == 0 {
		return
	}

	switch ba[0] {
	default:
		return

	case '+':
		if equalIgnoreCaseBytes(ba, infPlusShortBytes) || equalIgnoreCaseBytes(ba, infPlusLongBytes) {
			return math.Inf(1), true
		}
	case '-':
		if equalIgnoreCaseBytes(ba, infMinusShortBytes) || equalIgnoreCaseBytes(ba, infMinusLongBytes) {
			return math.Inf(-1), true
		}
	case 'n', 'N':
		if equalIgnoreCaseBytes(ba, nanBytes) {
			return math.NaN(), true
		}
	case 'i', 'I':
		if equalIgnoreCaseBytes(ba, infShortBytes) || equalIgnoreCaseBytes(ba, infLongBytes) {
			return math.Inf(1), true
		}
	}
	return
}
