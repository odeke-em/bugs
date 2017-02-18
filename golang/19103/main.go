package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	req, err := http.Get(`https://graph.facebook.com//v2.4/oauth/access_token`)
	if err != nil {
		panic(err.Error())
	}
	io.Copy(os.Stdout, req.Body)
	req.Body.Close()
}
