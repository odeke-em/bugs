package main

// #cgo LDFLAGS: -lverify
// #include </usr/local/lib/libverify.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func freeCString(cString *C.char) {
	C.free(unsafe.Pointer(cString))
}

func Init(serviceName string) {
	cServiceName := C.CString(serviceName)
	defer freeCString(cServiceName)
	key := C.CString("test")
	defer freeCString(key)

	cRes := C.Verify(cServiceName, key)
	defer freeCString(cRes)

	fmt.Println(C.GoString(cRes))
}

func main() {
	Init("my-service")
}
