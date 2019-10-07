package main

import (
    "flag"
    "fmt"
    "image/gif"
    "net/http"
    "os"
)

func exitIfError(err error) {
    if err == nil {
        return
    }

    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
}

func main() {
    gifURL := flag.String("url", "https://cdn.keycdn.com/img/analytics.gif", "the URL of the GIF")
    flag.Parse()
    res, err := http.Get(*gifURL)
    exitIfError(err)
    _, err = gif.DecodeAll(res.Body)
    _ = res.Body.Close()
    exitIfError(err)
}