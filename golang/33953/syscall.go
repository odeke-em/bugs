package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	log.SetFlags(0)

	f, err := os.Create("bx.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.Write([]byte("Aloha, World!")); err != nil {
		log.Fatal(err)
	}

	r0, err := fcntl(int(f.Fd()), syscall.F_FULLFSYNC, 0)
	log.Printf("r0: %#v\n", r0)
	log.Printf("e1: %T\n", err)
	log.Printf("e1: %t\n", err == nil)

	getFl, err := IsNonBlock(int(f.Fd()))
	log.Printf("nonBlocking: %x\nerr: %v", getFl, err)

        res, err := fcntl(int(f.Fd()), syscall.F_SETFL, getFl|syscall.O_NONBLOCK)
        log.Printf("Res: %#v, err: %#v\n", res, err)
	getFl2, err := IsNonBlock(int(f.Fd()))
        log.Printf("nonBlocking: %x\nerr: %v want: %x", getFl2, err,getFl|syscall.O_NONBLOCK)
}

func fcntl(fd int, cmd int, arg int) (int, error) {
	r, _, e := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	if e != 0 {
		return int(r), syscall.Errno(e)
	}
	return int(r), nil
}

func IsNonBlock(fd int) (int, error) {
	return fcntl(fd, syscall.F_GETFL, 0)
        /*
	if err != nil {
		return false, err
	}
	log.Printf("IsNonBlock: %x", flag)
	return flag&syscall.O_NONBLOCK != 0, nil
        */
}
