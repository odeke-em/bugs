package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	var stat unix.Stat_t
        if err := unix.Stat(os.Args[0], &stat); err != nil {
		log.Fatalf("Failed to stat: %v", err)
	}
	fmt.Printf("%+v\n", stat)
}
