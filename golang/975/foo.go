package foo

// #include <stdio.h>
import "C"

func PrintStuff() {
    C.printf(C.CString("Hello"))
}