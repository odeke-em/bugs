package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test26440(t *testing.T) {
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("test-%d", i)
		t.Run(name, func(tt *testing.T) {
			tt.Parallel()
			rnd := time.Duration(rand.Intn(537)) * time.Millisecond
			time.Sleep(rnd)
			if i&1 == 0 {
				t.Errorf("%s failed", name)
			}
		})
	}
}
