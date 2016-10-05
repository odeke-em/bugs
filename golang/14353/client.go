package main

import (
    "io"
    "log"
    "os"
    "net/http"
)

func main() {
     // works with "http://localhost:8080foobar/valid_path" too
    resp, err := http.Get("http://localhost:8080foobar/")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
        log.Fatal(err)
    }
}
