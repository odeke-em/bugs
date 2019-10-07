package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	m := new(sync.Map)
	key := "foo"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tick := time.NewTicker(1e3)
	// Writer routine
	go func() {
		i := 0
		for {
			m.Store(key, i)
			i += 1

			select {
			case <-tick.C:
				// Continue
			case <-ctx.Done():
				return
			}
		}
	}()

	// Delete routine
	go func() {
		i := 0
		for {
			m.Delete(key)
			i += 1

			select {
			case <-tick.C:
				// Continue
			case <-ctx.Done():
				return
			}
		}
	}()

	// Get routine for unexpungement
	go func() {
		i := 0
		for {
			m.Load(key)
			i += 1
			select {
			case <-tick.C:
				// Continue
			case <-ctx.Done():
				return
			}
		}
	}()

	for i := 0; i < 1e9; i++ {
		m.Range(func(key, value interface{}) bool {
			return true
		})
	}
}
