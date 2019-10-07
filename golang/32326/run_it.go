package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	nPipes := []int{0, 1, 5, 10, 20, 50, 75, 100}

	tmpDir, err := ioutil.TempDir("", "th")
	if err != nil {
		log.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	mainGoPath := filepath.Join(tmpDir, "main.go")
	if err := ioutil.WriteFile(mainGoPath, []byte(sourceCode), 0644); err != nil {
		log.Printf("writing main file %q: %v", mainGoPath, err)
		return
	}

	for _, n := range nPipes {
		if err := runIt(tmpDir, mainGoPath, n); err != nil {
			log.Printf("Error building for %d: %v\n", n, err)
		}
	}

	fz, err := os.Create("contents.zip")
	if err != nil {
		log.Printf("Failed to create contents.zip file: %v", err)
		return
	}
	defer fz.Close()

	zw := zip.NewWriter(fz)
	defer zw.Close()
	defer zw.Flush()
	err = filepath.Walk(tmpDir, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		zfh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		w, err := zw.CreateHeader(zfh)
		if err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(w, f)
		return err
	})
	if err != nil {
		log.Fatalf("filepath.Walk error: %v", err)
	}
}

func runIt(baseDir, mainGoPath string, n int) error {
	// Now run it and save it to the file.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", mainGoPath, "-dir", baseDir, "-n", fmt.Sprintf("%d", n))
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("exec error: %v\nOutput: %s\n", err, output)
	}
	return err
}

const sourceCode = `
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
	if n <= 0 {
	        return
        }
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
}`
