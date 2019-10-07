package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Rat)
	a.SetString("8.192")
	fmt.Println(a.FloatString(3))
	a.SetString("4.096")
	fmt.Println(a.FloatString(3))
	a.SetString("-213.09")
	fmt.Println(a.FloatString(3))
	a.SetString("92")
	fmt.Println(a.FloatString(3))
}
