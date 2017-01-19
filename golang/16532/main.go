package main

import (
    "log"
    "net/http"
)

func main() {
    _, err := http.Get("https://storage.googleapis.com/minikube/releases.json")
    if err != nil {
        log.Fatal(err)
    }
}