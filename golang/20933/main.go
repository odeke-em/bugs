package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

type closeWriter interface {
	CloseWrite() error
}

func main() {
	dataSentAfterHijack := "0123456789"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, connBuf, err := w.(http.Hijacker).Hijack()
		if err != nil {
			log.Fatalf("Failed to hijack backend connection: %s", err)
		}
		defer conn.Close()

		buf := &bytes.Buffer{}

		if err := r.Write(buf); err != nil {
			log.Fatalf("Failed to read request into buffer: %s", err)
		}

		fmt.Printf("Got HTTP Request:\n\n%s\n", buf.String())
		buf.Reset()

		// Introduce some delay before we begin reading more data from
		// the client. This simulates some operation that the server
		// might do to prepare resources. It also gives time for the
		// backgroundRead funtion to begin in the net/http lib.
		time.Sleep(100 * time.Millisecond)

		// Read more from the client. Include the connection buffer
		// which may contain data.
		// Note that this will read until EOF (when the client closes
		// its write-end of the connection).
		if _, err := io.Copy(buf, connBuf); err != nil {
			log.Fatalf("Failed to read more data from hijacked server connection: %s", err)
		}
		if buf.String() != dataSentAfterHijack {
			log.Fatalf("Expected to read %q after request, got %q", dataSentAfterHijack, buf.String())
		}

		fmt.Printf("Got data after Request: %q\n", buf.String())

		// Send that data back to the client.
		if _, err := io.Copy(conn, buf); err != nil {
			log.Fatalf("Failed to write data back to hijacked server connection: %s", err)
		}
	}))
	defer server.Close()

	conn, err := net.Dial("tcp", server.Listener.Addr().String())
	if err != nil {
		log.Fatalf("unable to connect to frontend: %s", err)
	}
	defer conn.Close()

	request, err := http.NewRequest("POST", server.URL, strings.NewReader("This is the requset body\n"))
	if err != nil {
		log.Fatalf("Failed to create requset: %s", err)
	}

	if err := request.Write(conn); err != nil {
		log.Fatalf("Failed to write request: %s", err)
	}

	// Wait for backgroundRead to begin.
	time.Sleep(100 * time.Millisecond)

	if _, err := io.Copy(conn, strings.NewReader(dataSentAfterHijack)); err != nil {
		log.Fatalf("Failed to write extra data to client connection: %s", err)
	}

	conn.(closeWriter).CloseWrite()

	// Should be able to read the same data back.
	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, conn); err != nil {
		log.Fatalf("Failed to read extra data back from client connection: %s", err)
	}

	if buf.String() != dataSentAfterHijack {
		log.Fatalf("Expected to read back %q, got %q", dataSentAfterHijack, buf.String())
	}
}
