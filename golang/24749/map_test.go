package map_test

import (
	"fmt"
	"strings"
	"testing"
)

type mode bool

const (
	reading = false
	writing = true
)

func mapAccess(m map[int]int, key int, mod mode) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch typ := r.(type) {
			case error:
				err = typ
			default:
				err = fmt.Errorf("%v", typ)
			}
		}
	}()

	if mod == reading {
		_, _ = m[key]
	} else {
		m[key] = 10
	}

	return
}

func TestMapAccess(t *testing.T) {
	mp := map[int]int{
		1: 2,
		2: 3,
		3: 4,
	}

	n := 6
	capturedErrorsChan := make(chan error)

	for i := 0; i < n; i++ {
		key := 3
		var mod mode = reading
		if (i & 1) != 2 {
			mod = writing
		}

		go func(key int, mod mode) {
			capturedErrorsChan <- mapAccess(mp, key, mod)
		}(key, mod)
	}

	isConcurrentReadWrite := func(err error) bool {
		return err != nil && strings.Contains(err.Error(), "concurrent map read and map write")
	}

	passCount := 0
	for i := 0; i < n; i++ {
		err := <-capturedErrorsChan
		if isConcurrentReadWrite(err) {
			passCount += 1
		}
	}

	t.Logf("passCount: %d out of %d\n", passCount, n)
}
