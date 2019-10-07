package main

import (
	"time"
)

func main() {
	var a *struct{ A int32 }
	println(helperCall(a.A))
}

func helperCall(a int32) time.Duration {
	return time.Duration(a)
}
