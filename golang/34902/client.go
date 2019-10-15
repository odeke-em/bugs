package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main() {
	var n int
	// 10 goroutines is a good number to see the panic happen quickly on
	// our test environments, though this may not apply to all machines
	flag.IntVar(&n, "n", 10, "number of goroutines")
	flag.Parse()

	// Request needs a larger body to see the error happen more quickly
	// It is reproducable with smaller body sizes, but it takes longer to fail
	bodyString := strings.Repeat("a", 2048)
	client := &http.Client{}

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				buf := bytes.NewBufferString(bodyString)
				req, _ := http.NewRequest("POST", "http://localhost:8080", buf)
				req.Header.Add("Expect", "100-continue")

				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("Request Failed: %s\n", err.Error())
				} else {
					resp.Body.Close()
				}
			}
		}()
	}
}
