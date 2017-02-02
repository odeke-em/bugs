package main

type m map[string]string

var _ = map[string]string{}       // ok
var _ = map[string]string{"": ""} // ok

func main() {
	var _ = map[string]string{}           // ok
	var _ = map[string]bool{}             // ok
	var _ = map[string]string{"a": "A"}   // error: cannot use _ as value
	var _ = map[string]bool{"true": true} // error: cannot use _ as value
}
