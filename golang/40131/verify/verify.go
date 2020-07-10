package main

// #include <stdlib.h>
import "C"

import (
	"encoding/hex"
	"strings"

	"./sec"
)

//export Verify
func Verify(serviceNameC *C.char, keyC *C.char) *C.char {
	serviceName := C.GoString(serviceNameC)

	key := []byte(C.GoString(keyC))
	sec.InitWithKey(key)

	res := "res"
	return C.CString(strings.Join([]string{serviceName, res, hex.EncodeToString(key)}, "-"))
}

func main() {

}
