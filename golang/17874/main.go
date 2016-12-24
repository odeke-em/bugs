package main

import (
	"fmt"

	"time"
)

func main() {

	p := fmt.Println
 	v, err := time.Parse("02.01.2006", "00.01.1970")
	p(v, err)


// expected result: "01.01.0001" plus an error 
// actual result:   1969-12-31 00:00:00 +0000 UTC <nil>
}
