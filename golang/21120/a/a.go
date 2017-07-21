package a

import "reflect"

type S struct {
	x int
}

func F() string {
	v := reflect.TypeOf(S{0})
	return v.PkgPath()
}