package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	signal.Ignore(syscall.SIGHUP)

	go func() {
		runtime.LockOSThread()
		pid := os.Getpid()
		tid := syscall.Gettid()
		for {
			syscall.Tgkill(pid, tid, syscall.SIGHUP)
		}
	}()

	c := make(chan os.Signal, 1)
	for {
		signal.Notify(c, syscall.SIGHUP)
		signal.Ignore(syscall.SIGHUP)
	}
}