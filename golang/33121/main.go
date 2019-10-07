package main

import (
	"debug/elf"
	"fmt"
	"strings"
)

func main() {
	data := "\u007fELF\x02\x01\x010000000000000" +
		"\x010000000000000000000" +
		"\x02\x00\x00\x00\x00\x00\x00\x0000000000\x00\x00\x00\x00" +
		"000\x0000"

	_, err := elf.NewFile(strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
}