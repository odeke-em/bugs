package main

import (
	"fmt"
        "runtime"
)

var foo = []byte{105, 57, 172, 152, 175, 139, 47, 108}

func main() {
        fmt.Println(runtime.Version())
	for i := 0; i < len(foo); i += 4 {
		fmt.Println(readLittleEndian32_1(foo[i], foo[i+1], foo[i+2], foo[i+3]))
		fmt.Println(readLittleEndian32_2(foo[i], foo[i+1], foo[i+2], foo[i+3]))
		fmt.Println(readLittleEndian32_3(foo[i], foo[i+1], foo[i+2], foo[i+3]))
                fmt.Println("")
	}
}

var x = false

func readLittleEndian32_1(a, b, c, d byte) uint32 {
	r := uint32(a) | (uint32(b) << 8) | (uint32(c) << 16) | (uint32(d) << 24)
	if !x {
		x = true
		fmt.Println(a, b, c, d, r)
	}
	return r
}

func readLittleEndian32_2(a, b, c, d byte) uint32 {
	return uint32(a) | (uint32(b) << 8) | (uint32(c) << 16) | (uint32(d) << 24)
}

func readLittleEndian32_3(a, b, c, d byte) uint32 {
	return (uint32(a) & 0xFF) | ((uint32(b) << 8) & 0xFF00) | ((uint32(c) << 16) & 0xFF0000) | ((uint32(d) << 24) & 0xFF000000)
}
