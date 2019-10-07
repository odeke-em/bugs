package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatalf("Failed to get an available port: %v", err)
	}

	defer ln.Close()
        ip, err := localIP(ln.Addr().String())
        log.Printf("ip: %s\nerr: %v\n", ip, err)

        return
}

func localIP(addr string) (ip net.IP, err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.TCPAddr)
	return localAddr.IP, nil
}
