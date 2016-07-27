package main

import (
    "reflect"
)

type S struct {}

func (s *S) Get() int { return 0 }

type I interface {
    F(*S)
}

func main() {
    ft := reflect.TypeOf((*I)(nil)).Elem()
    m := ft.Method(0)
    mt := m.Type
    for i := 0; i < mt.NumIn(); i++ {
        if mt.In(i).PkgPath() != "" {
            panic(mt.In(i).PkgPath())
        }
    }
}