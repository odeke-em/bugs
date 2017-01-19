// plugin.go

package main

import "C"
import "net/http"

func F() {
	_, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
}