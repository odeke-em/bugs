package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var PathSeparator = fmt.Sprintf("%c", os.PathSeparator)

func leastNonExistantRootPatched(p string) string {
    last := ""
    for p != "" {
        fInfo, _ := os.Stat(p)
        if fInfo != nil {
            break
        }
        last = p
        p, _ = filepath.Split(strings.TrimRight(p, PathSeparator))
        fmt.Println("Still spinning: ", p)
    }
    return last
}

func leastNonExistantRootInfinite(p string) string {
    last := ""
    for p != "" {
        fInfo, _ := os.Stat(p)
        if fInfo != nil {
            break
        }
        last = p
        p, _ = filepath.Split(p)
        fmt.Println("Still spinning: ", p)
    }
    return last
}

func main() {
    leastNonExistantRootPatched("/mnt/ghosting/symver/projects/gophers/alophering/change.go")
    leastNonExistantRootInfinite("/mnt/ghosting/symver/projects/gophers/alophering/change.go/")
}
