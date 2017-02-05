package issue_test

import (
	"io"
	"net"
	. "net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Test that a hanging Request.Body.Read from another goroutine can't
// cause the Handler goroutine's Request.Body.Close to block.
// See golang.org/issue/7121
func TestRequestBodyCloseDoesntBlock(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in -short mode")
	}
	// defer afterTest(t)

	readErrCh := make(chan error, 1)
	errCh := make(chan error, 2)

	server := httptest.NewServer(HandlerFunc(func(rw ResponseWriter, req *Request) {
		go func(body io.Reader) {
			_, err := body.Read(make([]byte, 100))
			readErrCh <- err
		}(req.Body)
		time.Sleep(500 * time.Millisecond)
	}))
	defer server.Close()

	closeConn := make(chan bool)
	defer close(closeConn)
	go func() {
		conn, err := net.Dial("tcp", server.Listener.Addr().String())
		if err != nil {
			errCh <- err
			return
		}
		defer conn.Close()
		_, err = conn.Write([]byte("POST / HTTP/1.1\r\nConnection: close\r\nHost: foo\r\nContent-Length: 100000\r\n\r\n"))
		if err != nil {
			errCh <- err
			return
		}
		// And now just block, making the server block on our
		// 100000 bytes of body that will never arrive.
		<-closeConn
	}()

	select {
	case err := <-readErrCh:
		t.Logf("Read error = %v", err)
	case err := <-errCh:
		t.Error(err)
	case <-time.After(5 * time.Second):
		t.Error("timeout")
	}
}
