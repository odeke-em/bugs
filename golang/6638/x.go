package p

import "unsafe"

type T interface {
    m() [unsafe.Sizeof(T(nil).m())]int
}
