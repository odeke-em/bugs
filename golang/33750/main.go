package main

import (
	"fmt"
	"strconv"
)

func main() {
	d, err := strconv.ParseInt("0X1_2350_0_0e10", 0, 64)
	fmt.Printf("d: %v\nerr: %v\n", d, err)

        f, err := strconv.ParseFloat("0x1_2.3_4_5e1_12", 64)
        fmt.Printf("f: %v\nerr: %v\n", f, err)
}
