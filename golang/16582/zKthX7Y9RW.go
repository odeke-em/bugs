package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	client := &http.Client{}

	errs := make(chan error)
	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Get("https://www.google.com/")
			if err != nil {
				errs <- err
				return
			}
			resp.Body.Close()
		}()
	}
	go func() {
		wg.Wait()
		close(errs)
	}()

	fail := false
	for err := range errs {
		fmt.Fprintln(os.Stderr, err)
	}

	if fail {
		os.Exit(1)
	}
}
