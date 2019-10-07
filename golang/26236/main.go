package main

import (
	"path"
)

func main() {
        codeDir := "foo"
	file1 := path.Join(codeDir, "go.mod")
	gomod1, err1 := foo(file1)
	gomod2, err2 := foo(file2)
	found1 := err1 == nil && true
	found1 := err2 == nil && true
	if err2 == nil && !found2 {
	}
	if found1 && found2 {
	}
        if gomod != "" {
        }
	/*
		f1 := func() int {
			return p2
		}
		_ = file2
		p1, x := f1(), y
	*/
}

func foo(string) (string, error) {
	return "", nil
}
