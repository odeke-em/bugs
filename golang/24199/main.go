package main

import "syscall"

func main() {
	go func() {
		for {
			sig := <-sc
			switch sig {
			case syscall.SIGINT:
				//proc.Signal(syscall.SIGINT)
			default:
				proc.Signal(sig)
			}
		}
	}()
}
