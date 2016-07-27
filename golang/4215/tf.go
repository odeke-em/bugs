// errorcheck

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type node struct {
	left, right *node
	data        interface{}
	visitCount  uint64
}

func foo() (int, int) {
	return 2.3 // ERROR "not enough arguments to return, got (untyped number) want (int, int)"
}

func foo2() {
	return int(2), 2 // ERROR "too many arguments to return, got (int,untyped number) want ()"
}

func foo3() (a, b, c, d int) {
	return 1 // ERROR "not enough arguments to return, got (untyped number) want (int, int, int, int)"
}

func (n *node) walk() {
	if n == nil {
		return 0 // ERROR "too many arguments to return, got (untyped number)"
	}
	n.left.walk()
	n.right.walk()
	n.visitCount += 1
	return n.visitCount // ERROR too many arguments to return, got (uint64) want ()
} // ERROR "k
