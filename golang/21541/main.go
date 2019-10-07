package main

import (
    "fmt"
    "os"
    "syscall"
)

const (
    anyFile = "any_file.txt"
)

func main() {
    var creationTime syscall.Filetime
    var lastAccessTime syscall.Filetime
    var lastWriteTime syscall.Filetime

    utf16FileName, err := syscall.UTF16FromString(anyFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to convert string %q to UTF16: %v", anyFile, err)
    }

    err = syscall.GetFileTime(utf16FileName, &creationTime, &lastAccessTime, &lastWriteTime)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting FileTime of file %v: %v", anyFile, err)
        return
    }
    // do something with creationTime, lastAccessTime and/or lastWriteTime
    return
}