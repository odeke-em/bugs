package main

import (
    "os"
    "io/ioutil"
    "path"
    "fmt"
    "os/exec"
)

func main() {
    tmpDir, _ := ioutil.TempDir("", "")
    defer os.RemoveAll(tmpDir)

    realDir := path.Join(tmpDir, "real")
    _ = os.MkdirAll(realDir, 0755)

    symlinkDir := path.Join(tmpDir, "symlink")
    _ = os.Symlink(realDir, symlinkDir)

    cmd := exec.Command("pwd")
    cmd.Dir = symlinkDir
    stdout, _ := cmd.Output()
    fmt.Printf("Expected: %v\nActual:   %v", symlinkDir, string(stdout))
}
