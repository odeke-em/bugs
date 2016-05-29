package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func reproTrial(path string, content []byte) error {
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = os.RemoveAll(path)
	}()

	retrieved, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if !bytes.Equal(retrieved, content) {
		return fmt.Errorf("file %q content doesn't match, read %s expected %s\n", path, retrieved, content)
	}
	return nil
}

func main() {
	outfilePath := "temp"
	content := []byte("foo")

	if err := reproTrial(outfilePath, content); err != nil {
		log.Fatal(err)
	}
}
