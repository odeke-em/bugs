package main

import (
    "fmt"
    "net"
)

func main() {
    addr, err := net.LookupHost("storage.googleapis.com")
    fmt.Println(addr, err)
}