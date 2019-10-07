package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	tmpDir, err := ioutil.TempDir("", "30520")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	evalDir, err := filepath.EvalSymlinks(tmpDir)
	if err != nil {
		log.Fatalf("EvalSymlinks: %v", err)
	}
	log.Printf("EvalDir: %s\nTmpDir:  %s", evalDir, tmpDir)
}
