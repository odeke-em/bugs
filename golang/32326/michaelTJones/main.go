
package main

import (
	"flag"
        "fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/trace"
)

func osPipesIO(n int) {
	r := make([]*os.File, n)
	w := make([]*os.File, n)
	for i := 0; i < n; i++ {
		rp, wp, err := os.Pipe()
		if err != nil {
			for j := 0; j < i; j++ {
				r[j].Close()
				w[j].Close()
			}
			log.Fatal(err)
		}
		r[i] = rp
		w[i] = wp
	}

	creading := make(chan bool, n)
	cdone := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func(i int) {
			var b [1]byte
			creading <- true
			if _, err := r[i].Read(b[:]); err != nil {
				log.Printf("r[%d].Read: %v", i, err)
			}
			if err := r[i].Close(); err != nil {
				log.Printf("r[%d].Close: %v", i, err)
			}
			cdone <- true
		}(i)
	}
	for i := 0; i < n; i++ {
		<-creading
	}

	for i := 0; i < n; i++ {
		if _, err := w[i].Write([]byte{0}); err != nil {
			log.Printf("w[%d].Read: %v", i, err)
		}
		if err := w[i].Close(); err != nil {
			log.Printf("w[%d].Close: %v", i, err)
		}
		<-cdone
	}
}

func main() {
        baseDir := flag.String("dir", "", "the base directory to place execution traces")
        n := flag.Int("n", 0, "the number of *os.Pipe to create")
        flag.Parse()
	f, err := os.Create(filepath.Join(*baseDir, fmt.Sprintf("trace-%d.txt", *n)))
	if err != nil {
		log.Fatalf("Failed to create trace file: %v", err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	osPipesIO(*n)
}
