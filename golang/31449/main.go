package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to grab an available socket: %v", err)
	}
	defer ln.Close()

	pr, pw, err := os.Pipe()
	if err != nil {
		log.Fatalf("Failed to get a named pipe: %v", err)
	}
	defer pr.Close()
	defer pw.Close()
	log.Printf("pr.Fd=%d pw.Fd=%d", pr.Fd(), pw.Fd())

	tcpLn := ln.(*net.TCPListener)
	
	go func() {
		for {
			conn, err := ln.Accept()
			if false {
				src, err := tcpLn.SyscallConn()
				if err != nil {
					log.Fatalf("Failed to get the associated file: %v", err)
				}
				src.Control(func(fd uintptr) {
					log.Printf("Invoked control on: %d\n", fd)
				})
			}
			if err != nil {
				log.Fatalf("Failed to accept a client: %v", err)
			}

                        f, err := os.Create("foo.txt")
                        if err != nil {
                            log.Fatalf("Failed to create blank file: %v", err)
                        }

                        if false {
                        tcpConn := conn.(*net.TCPConn)
                        tcpConnFile, err := tcpConn.File()
                        if err != nil {
                            log.Fatalf("Failed to obtain file for TCPConn: %v", err)
                        }
                        log.Printf("tcpConnFile.Fd = %d\n", tcpConnFile.Fd())
			// At this point, we've successfully accepted the connection.
			if err := syscall.Dup2(int(tcpConnFile.Fd()), int(f.Fd())); err != nil {
				log.Fatalf("Failed to dup2: %v", err)
			}
                    }
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\nConnection: close\r\nTransfer-Encoding: chunked\r\n\r\n4\r\nOKAY\r\n0\r\n\r\n")
			conn.Close()
		}
	}()
	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

		tcn := conn.(*net.TCPConn)
		tcn.SetLinger(0)
		tcn.SetKeepAlivePeriod(1 * time.Second)

                        tcnFile, err := tcn.File()
                        if err != nil {
                            log.Fatalf("Failed to obtain file for TCPConn: %v", err)
                        }
                        log.Printf("tcpCnFile.Fd = %d\n", tcnFile.Fd())

	// We expect back a response.
	blob, err := ioutil.ReadAll(conn)
	if err != nil {
		oe := err.(*net.OpError)
		log.Printf("IsTemporary: %v IsTimeout: %v\n", oe.Temporary(), oe.Timeout())
		log.Fatalf("Failed to read response from server: %v, %w", err, oe)
	}
	log.Printf("Blob: %s\n", blob)
}
