package main_test

import (
	"runtime/debug"
	"testing"
	"time"
)

func TestRaceGoroutineID(t *testing.T) {
	var a int
	f := func() {
		debug.PrintStack()
		<-time.After(time.Second)
		debug.PrintStack()
		for i := 0; i < 10000000; i++ {
			a = 5
		}
	}
	shortWork := func() {
		_ = 10
		if 8&2 != 0 {
			panic("Foo bar")
		}
	}
	go shortWork()
	go f()
	f()
	t.Log(a)
}
