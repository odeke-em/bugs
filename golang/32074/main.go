package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	pr, pwc := io.Pipe()
	go func() {
		defer pwc.Close()
		for i := 0; i < 10; i++ {
			io.Copy(pwc, strings.NewReader(strings.Repeat("a", 20)))
		}
	}()

	req, _ := http.NewRequest("POST", "https://example.org/foo", pr)
	req.TransferEncoding = []string{"chunked"}
	blob, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s\n", blob)
}
