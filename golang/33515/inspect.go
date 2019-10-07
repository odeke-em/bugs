package main

import (
	"fmt"
	"log"
	"net"

        "golang.org/x/sys/unix"
)

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to grab an available address: %v", err)
	}
	defer ln.Close()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		log.Fatalf("Failed to dial to listener: %v", err)
	}
	defer conn.Close()

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatalf("Unfortunately never made a TCP connection: %T", conn)
	}
	fmt.Printf("Addr: %s\n", ln.Addr())

        file, err := tcpConn.File()
        if err != nil {
            log.Fatalf("Failed to find the associated file: %v", err)
        }
        fd := file.Fd()

        tcpInfo, err := unix.GetsockoptTCPInfo(int(fd), unix.SOL_TCP, unix.TCP_INFO)
        if err != nil {
            log.Fatalf("Failed to get socket option: %v", err)
        }
        log.Printf("TCPInfo: %#v\n", tcpInfo)
}
