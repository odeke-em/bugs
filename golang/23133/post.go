package main

import (
	"fmt"
	"net/http"
)

func Init() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Init Error:", err)
	}
	res.Body.Close()
}