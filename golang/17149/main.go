package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    execCmd("./folder name/script.sh")
    execCmd("./folder name/script.sh", "one param")
}

func execCmd(path string, args ...string) {
    fmt.Printf("Running: %q %q\n", path, strings.Join(args, " "))
    cmd := exec.Command(path, args...)
    bs, err := cmd.CombinedOutput()
    fmt.Printf("Output: %s", bs)
    fmt.Printf("Error: %v\n\n", err)
}
