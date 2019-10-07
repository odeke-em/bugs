package main

import (
        "fmt"
        "io/ioutil"
        "os"
        "os/exec"

        "golang.org/x/sys/unix"
)

func main() {
        dir, err := ioutil.TempDir("", "fd-escape")
        if err != nil {
                panic(err)
        }
        defer os.RemoveAll(dir)

        f, err := os.Open(dir)
        if err != nil {
                panic(err)
        }

        // copy from os.RemoveAll go 1.12 openFdAt without O_CLOEXEC
        if _, err := unix.Openat(int(f.Fd()), "/tmp", os.O_RDONLY, 0); err != nil {
                panic(err)
        }

        cmd := exec.Command("sleep", "10")
        if err := cmd.Start(); err != nil {
                panic(err)
        }

        fds, err := ioutil.ReadDir(fmt.Sprintf("/proc/%d/fd", cmd.Process.Pid))
        if err != nil {
                panic(err)
        }

        for _, fd := range fds {
                fname, err := os.Readlink(fmt.Sprintf("/proc/%d/fd/%s", cmd.Process.Pid, fd.Name()))
                if err != nil {
                        panic(err)
                }
                fmt.Println(fd.Name(), " --> ", fname)
        }
        cmd.Wait()
}
