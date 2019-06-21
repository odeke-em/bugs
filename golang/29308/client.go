package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"syscall"
	"time"
)

const url = "{{.URL}}"

var maxIdleTimeout time.Duration

func init() {
	idleTimeout, err := time.ParseDuration("{{.MaxIdleTimeout}}")
	if err != nil {
		log.Fatalf("Failed to parse MaxIdleTimeout: %v", err)
	}
	maxIdleTimeout = idleTimeout
}

func main() {
	pid := os.Getpid()
	ownProc, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalf("Could find own process by pid:%d %v", pid, err)
	}

	tr := &http.Transport{
		IdleConnTimeout: maxIdleTimeout,
	}
        tr.IdleConnTimeout = 0
	client := &http.Client{
		Transport: tr,
	}

	// Warm-up phase.
	paths := []string{
		"/initial",
		"/lag",
	}

	var stopOnce sync.Once

	for {
		for _, path := range paths {
			purl := url + path
			log.Printf("\npurl: %s\n", purl)
			res, err := client.Get(purl)
			if err != nil {
				log.Fatalf("Failed to make request(%q): %v", purl, err)
			}
			if _, err := io.Copy(os.Stdout, res.Body); err != nil {
				log.Fatalf("%q failed to read body: %v", purl, err)
			}

			if false {
				stopOnce.Do(func() {
					log.Println("About to send SIGTSTP to ourselves")
					// Send SIGTSTP to ourselves.
					ownProc.Signal(syscall.SIGTSTP)
				})
			}
			time.Sleep(maxIdleTimeout / time.Duration(len(paths)))
		}
		println("\n")
	}
}
