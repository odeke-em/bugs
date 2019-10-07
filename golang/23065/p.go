// errorcheck

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

type fancyMap map[string]string

func foo(fancyMap fancyMap) {
	var _ fancyMap // ERROR "fancyMap is a parameter that shadows the type"
}

func foo2(fancyMap fancyMap) {
	_ = make(fancyMap) // ERROR "fancyMap is a parameter that shadows the type"
}
