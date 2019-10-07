package main

import "unsafe"

func f() unsafe.Pointer
func g(u uintptr) int {
	return int(u)
}

// Scenario 1.
//  tmp := f()
//  return g(uintptr(tmp))
//  runtime.KeepAlive(tmp)
func a() int {
	return g(uintptr(f()))
}

// Scenario 2.
//  tmp := f()
//  return h(g(uintptr(tmp)))
//  runtime.KeepAlive(tmp)
func b() {
	h(g(uintptr(f())))
}

func h(i int) {
	_ = i
}
