package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	f1, err := ioutil.TempFile("", "box")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(f1.Name())
	t1 := time.Date(2004, time.September, 26, 16, 15, 14, 1921999, time.UTC)
	if err := os.Chtimes(f1.Name(), t1, t1); err != nil {
		log.Fatal(err)
	}

	fi, err := os.Stat(f1.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("want: %v, fi: %+v\n", t1, fi.ModTime())
}
