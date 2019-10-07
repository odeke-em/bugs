package main

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
)

func ioctl(fd, cmd, ptr uintptr) error {
	if _, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr); e != 0 {
		return e
	}
	return nil
}

func ptsname(fd uintptr) (string, error) {
	n := make([]byte, 128)
	if err := ioctl(fd, syscall.TIOCPTYGNAME, uintptr(unsafe.Pointer(&n[0]))); err != nil {
		return "", err
	}
	for i, c := range n {
		if c == 0 {
			return string(n[:i]), nil
		}
	}
	return "", errors.New("TIOCPTYGNAME string not NUL-terminated")
}

func main() {
	// Open master.
	ptmx, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer func() { _ = ptmx.Close() }()
	fd := ptmx.Fd()

	// Fetch next available slave.
	sname, err := ptsname(fd)
	if err != nil {
		panic(err)
	}
	// Grant/unlock slave.
	if err := ioctl(fd, syscall.TIOCPTYGRANT, 0); err != nil {
		panic(err)
	}
	if err := ioctl(fd, syscall.TIOCPTYUNLK, 0); err != nil {
		panic(err)
	}

	// Open slave.
	t, err := os.OpenFile(sname, os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer func() { _ = t.Close() }()

	// Read master.
	if _, err := ptmx.Read([]byte{0}); err != nil {
		panic(err)
	}
}
