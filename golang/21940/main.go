package main

import "time"

func init() {
	type T struct {
		x int
	}
	var t = &T{}

	apiCalls := func(start int, m *T) {
		for i := start; ; i++ {
			m.x = i
			<-time.After(10 * time.Millisecond)
		}
	}

	for i := 0; i < 10; i++ {
		go apiCalls(i, t)
	}
}

func main() {
	for {
		<-time.After(11 * time.Millisecond)
	}
}
