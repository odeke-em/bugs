package main

import (
    "fmt"
    "os/user"
)

func main() {
    u, err := user.Lookup("emmanuelodeke")
    if err != nil {
        fmt.Printf("%s", err)
        return
    }

    fmt.Printf("%V", u)
}
