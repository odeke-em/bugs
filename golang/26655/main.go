package main

// #cgo CFLAGS: -g
// #include <stdlib.h>
// #include "cLogic.h"
import "C"
import (
    "fmt"
    "unsafe"
)    

func main() {
    myString := "DUMMY"
    cMyString := C.CString(myString)
    defer C.free(unsafe.Pointer(cMyString))

    cMyInt := C.int(10)

    cResult := C.MyCFunction(cMyString, cMyInt) // Result is type *_Ctype_schar (int8_t *)
    goResult := C.GoString(cResult)
    fmt.Println("GoResult: " + goResult + "\n")
}