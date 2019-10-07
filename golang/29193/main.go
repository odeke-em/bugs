package main

import (
    "net/http"
    _ "net/http/httptest"
)

func main() {
    cst := httptest.NewServer(http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request) {

    }))
    defer cst.Close()
}
