package main

/*
#include <stdio.h>
typedef unsigned long long int ulong_long;
ulong_long GCD(ulong_long u, ulong_long v) {
  ulong_long rem = 0;
  while (v) {
    rem = u % v;
    u = v;
    v = rem;
  }

  return u;
}
*/
import "C"

import (
	"fmt"
)

func GCD(u, v int64) int64 {
	var res C.ulong_long = C.GCD(C.ulong_long(u), C.ulong_long(v))
	return int64(res)
}

func main() {
	v := GCD(10, 2)
	fmt.Printf("v=%v\n", v)
}
