package key

import "unsafe"

//go:linkname strhash runtime.strhash
func strhash(a unsafe.Pointer, h uintptr) uintptr

func hash(s string) uintptr {
    return strhash(unsafe.Pointer(&s), 0)
}