package main

import "time"

func f() {
}

func main() {
	f := func(p int) (l int64) {
		x := 10
		y := int64(x * 12)
		if y*time.Now().Unix()%2 == 0 {
			return 12
		}
		return y
	}
	if f(5)%2 == 0 {
		println("Produced even")
	}
	println("Foo")
}
