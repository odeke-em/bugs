package main

/*
#include <signal.h>
#include <unistd.h>

void foo() {
    usleep(17000000);
    raise(11);
}
*/
import "C"

import (
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	doneCh := make(chan bool)
	n := 10
	for i := 0; i < n; i++ {
		go func() {
			var b [2001]byte
			if false {
				b[1023] = 'b'
			}
			<-time.After(1 * time.Minute)
			doneCh <- true
		}()
	}

	fin := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-fin:
				return
			case <-time.After(time.Duration(rng.Intn(377)) * time.Millisecond):
				runtime.GC()
			}
		}
	}()

	go func() {
		/*
		   defer func() {
		       var f *int
		       if r := recover(); r != nil {
		           _ = *f
		       }
		   }()
		*/
		C.foo()
	}()

	for i := 0; i < n; i++ {
		<-doneCh
	}

	close(fin)
}
