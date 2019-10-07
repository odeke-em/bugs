package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var negative  = ((^0)<<24)
    sz := uintptr(negative)
    sp := sz>>24
    fmt.Printf("%x uinptr.size: %d\n", sp, unsafe.Sizeof(uintptr(0)))
}
