package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	var flush bool
	flag.BoolVar(&flush, "flush", false, "if set, will flush the buffered io before exiting")
	flag.Parse()

	br := bufio.NewWriter(os.Stdout)
	logger := log.New(br, "", log.Ldate)
	logger.Printf("%s\n", strings.Repeat("This is a test\n", 5))
	if flush {
		br.Flush()
	}
	logger.Fatalf("exiting now!")
}
