package main

import (
        "log"
        "net/http"
        "os"
)

func main() {
        resp, err := http.Get("https://golang.org:")
        if err != nil {
                log.Printf("err: %v", err)
                os.Exit(1)
        }
        log.Printf("resp: %v", resp.Request)
}
