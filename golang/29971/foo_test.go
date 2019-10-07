package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestPermanentlyBrokenConnection(t *testing.T) {
	cst := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	defer cst.Close()

	cstU, _ := url.Parse(cst.URL)
	conn, err := tls.Dial("tcp", cstU.Host, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		t.Fatalf("Dial err: %v", err)
	}

	if err = conn.SetDeadline(time.Now()); err != nil {
		t.Fatalf("conn.SetDeadline(time.Now()): %v", err)
	}

	if _, err = conn.Write([]byte("should fail")); err == nil {
		t.Fatal("conn.Write succeeded when it should have failed")
	}

	// Clear deadline
	err = conn.SetDeadline(time.Time{})
	if err != nil {
		t.Fatalf("final conn.SetDeadline(time.Time{}): %v", err)
	}

	_, err = conn.Write([]byte("This connection is permanently broken"))
	if err != nil {
		fmt.Println("Write err", err)
	}

	ne := err.(net.Error)
	if ne.Temporary() {
		t.Fatal("Got a net.Error that apparently is temporary")
	}
}
