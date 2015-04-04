// Author: @sykesm

package main

import (
    "fmt"
    "io"
    "syscall"
)

import "os/exec"

func main() {
    pr, pw := io.Pipe()

    cmd := exec.Command("/bin/bash", "-c", "cat")
    cmd.Stdin = pr

    started := make(chan error)
    done := make(chan error)
    go func() {
        started <- cmd.Start()
        done <- cmd.Wait()
    }()

    <-started
    println("Started")
    pw.Write([]byte("hello"))

    cmd.Process.Signal(syscall.SIGTERM)

    err := <-done
    println("Done")
    if err != nil {
        fmt.Printf("wait err: %s\n", err.Error())
    }
}
