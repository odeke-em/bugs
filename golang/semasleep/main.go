//+build !windows

package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	tempDir, err := ioutil.TempDir("", "signal-refresh")
	if err != nil {
		log.Fatalf("Failed to create the temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Given the simple program below
	repro := `
    package main

    import "time"

    func main() {
        <-time.After(2 * time.Second)
    }
    `
	mainPath := filepath.Join(tempDir, "main.go")
	if err := ioutil.WriteFile(mainPath, []byte(repro), 0755); err != nil {
		log.Printf("Failed to create temp file for repro.go: %v", err)
		return
	}
	binaryPath := filepath.Join(tempDir, "binary")

	out, err := exec.Command("go", "build", "-o", binaryPath, mainPath).CombinedOutput()
	if err != nil {
		log.Printf("Failed to compile the binary: err: %v\nOutput: %s\n", err, out)
		return
	}
	if err := os.Chmod(binaryPath, 0755); err != nil {
		log.Printf("Failed to chmod binary: %v", err)
		return
	}

	// Now run the binary
	cmd := exec.Command(binaryPath)
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start command: %v", err)
		return
	}
	doneCh := make(chan error)
	go func() {
		doneCh <- cmd.Wait()
	}()

	// Now that we've started the repro, we
	// can continuously send to it signal SIGIO.
	unfixedTimer := time.NewTimer(4 * time.Second)
	for {
		select {
		case <-time.After(600 * time.Millisecond):
			// Send the pesky signal that toggle spinning
			// till infinity if #27520 is not fixed!!
			cmd.Process.Signal(syscall.SIGIO)
			log.Println("Sent SIGIO")

		case <-unfixedTimer.C:
			log.Println("Unfortunately the issue hasn't yet been fixed!")
			cmd.Process.Signal(syscall.SIGKILL)
			return

		case err := <-doneCh:
			if err != nil {
				log.Printf("The program returned but unfortunately with an error: %v", err)
			} else {
				log.Println("Hooray, the issue is fixed!!")
			}
			return
		}
	}
}
