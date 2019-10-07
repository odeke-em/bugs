package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var mainSourceCode = []byte(`
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	defer fmt.Printf("Spent: %s\n", time.Since(time.Now()))

	var b [1]byte
	_, err := os.Stdin.Read(b[:])
	if err != nil {
		log.Fatalf("Failed to read a byte from stdin: %v", err)
	}
}`)

func main() {
	log.SetFlags(0)

	tmpDir, err := ioutil.TempDir("", "killproc")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	mainFile := filepath.Join(tmpDir, "main.go")
	if err := ioutil.WriteFile(mainFile, mainSourceCode, 0644); err != nil {
		log.Fatalf("Failed to write source code to main file: %v", err)
	}

	// 1. Build the binary from the source code.
	binaryPath := filepath.Join(tmpDir, "killproc")
	output, err := exec.Command("go", "build", "-o", binaryPath, mainFile).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to build the binary: %v;\nOutput: %s", err, output)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 2. Now that we have the binary, run it.
	binCmd := exec.CommandContext(ctx, binaryPath)
	stdin := new(bytes.Buffer)
	binCmd.Stdin = stdin
	stdoutStderr := new(bytes.Buffer)
	binCmd.Stdout = stdoutStderr
	binCmd.Stderr = stdoutStderr
	if err := binCmd.Start(); err != nil {
		log.Fatalf("Failed to run the binary: %v", err)
	}

	binPid := binCmd.Process.Pid
	// 3. Look for the process in the existent case.
	foundProc, err := os.FindProcess(binPid)
	if err != nil {
		log.Fatalf("Failed to find the process right after starting: %v", err)
	}
	if foundProc == nil {
		log.Fatal("Oddly the found process is nil!")
	}

	// 3.1. Send it that single byte so that it completes.
	stdin.WriteString("Hello, World!")

	// 3.2. Wait for the process to terminate.
	if err := binCmd.Wait(); err != nil {
		log.Fatalf("Failed to end the process: %v", err)
	}

	// Explicitly even cancel the process to ensure it has terminated.
	cancel()

	// By this point we are sure the process has
	// ended and can now examine if it still lingers.
	log.Printf("Output from process: %s\n", stdoutStderr.String())

	// 4. Now search again from the process.
	foundAgainProc, err := os.FindProcess(binPid)
	if err != nil {
		// We are safe! Test passes.
		return
	}

	if foundAgainProc == nil {
		log.Fatal("Oddly err == nil, yet foundAgainProc is nil too")
	}

	switch goos := runtime.GOOS; goos {
	case "windows":
		log.Fatal("Oops issue https://golang.org/issue/33814 still lingers!")

	default:
		// On UNIXes, it is documented as per https://golang.org/pkg/os/#FindProcess
		// that findProcess is a noop.
		log.Printf("On GOOS=%s, os.FindProcess returned a non-blank process, please examine https://golang.org/pkg/os/#FindProcess", goos)
	}
}
