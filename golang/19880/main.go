package main

type T struct {
	f [1]int
}

func a() {
	_ = T
}

func b() {
	var v [len(T{}.f)]int
	_ = v
}
