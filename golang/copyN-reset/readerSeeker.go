package main

import (
  "io"
  "fmt"
)

type RS struct {
}

var _ io.ReadSeeker = &RS{}

func (rs *RS) Read(b []byte) (int, error) {
	return 0, fmt.Errorf("unimplemented")
}

func (rs *RS) Seek(offset int64, whence int) (int64, error) {
	return 0, fmt.Errorf("unimplemented")
}
