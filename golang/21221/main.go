package main

import (
	"unsafe"
)

func main() {
const v = unsafe.Pointer(uintptr(1))
	a1 := unsafe.Pointer(uintptr(0))
	a2 := unsafe.Pointer(nil)
	println(a1 == unsafe.Pointer(nil))
	println(a2 == unsafe.Pointer(uintptr(0)))
	println(a1 == a2)
}
