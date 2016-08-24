package main

import (
	"fmt"
)

func main() {
	mf := [4]string{}
	mf[0] = "hey"

	strs := make([]string, 4)
	fmt.Printf("len: %d cap: %d\n", len(strs), cap(strs))
	strs[0] = "user"
	strs[1] = "alex"
	strs[2] = "password"
	strs[3] = "foo"
	strs["4"] = "database"
	strs[5] = "ar_test"
	fmt.Println(strs)
}
