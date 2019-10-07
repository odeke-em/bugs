package main

/*
typedef long double             SmiFloat128;

typedef struct SV {
    union {
        SmiFloat128         float128;
    } value;
} SV;
*/
import "C"

type ts struct {
	tv *C.SV
}

func main() {}
