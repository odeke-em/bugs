package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type loopbackAddr string
type loopbackConn struct {
	addr loopbackAddr
	r    *io.PipeReader
	w    *io.PipeWriter
}

func (a loopbackAddr) Network() string { return "loopback" }
func (a loopbackAddr) String() string  { return string(a) }

func (c *loopbackConn) Read(b []byte) (n int, err error)   { return c.r.Read(b) }
func (c *loopbackConn) Write(b []byte) (n int, err error)  { return c.w.Write(b) }
func (c *loopbackConn) Close() error                       { c.r.Close(); return c.w.Close() }
func (c *loopbackConn) LocalAddr() net.Addr                { return c.addr }
func (c *loopbackConn) RemoteAddr() net.Addr               { return c.addr }
func (c *loopbackConn) SetDeadline(t time.Time) error      { return nil }
func (c *loopbackConn) SetReadDeadline(t time.Time) error  { return nil }
func (c *loopbackConn) SetWriteDeadline(t time.Time) error { return nil }

type loopbackListener struct {
	addr  loopbackAddr
	ch    chan *loopbackConn
	close chan bool
}

func (l *loopbackListener) Accept() (net.Conn, error) {
	conn, ok := <-l.ch
	if !ok {
		return nil, io.ErrClosedPipe
	}
	return conn, nil
}

func (l *loopbackListener) Close() error   { close(l.close); return nil }
func (l *loopbackListener) Addr() net.Addr { return l.addr }

func (l *loopbackListener) dial() (net.Conn, error) {
	lr, rw := io.Pipe()
	rr, lw := io.Pipe()
	localConn := &loopbackConn{addr: l.addr, r: lr, w: lw}
	remoteConn := &loopbackConn{addr: l.addr, r: rr, w: rw}
	select {
	case <-l.close:
		return nil, io.ErrClosedPipe
	case l.ch <- remoteConn:
		return localConn, nil
	}
}

type dialFunc func() (net.Conn, error)

var addrSeq = 0

func NewLoopbackAdapter() (net.Listener, dialFunc) {
	addr := addrSeq
	addrSeq++
	l := &loopbackListener{
		addr:  loopbackAddr(fmt.Sprintf("%d", addr)),
		ch:    make(chan *loopbackConn),
		close: make(chan bool),
	}
	return l, l.dial
}

func handler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	return
}

func serve(c net.Conn) {
	defer c.Close()
	s := &http2.Server{}
	s.ServeConn(c, &http2.ServeConnOpts{Handler: http.HandlerFunc(handler)})
}

func main() {
	l, dialer := NewLoopbackAdapter()
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				panic(err)
			}
			go serve(conn)
		}
	}()

	t := &http2.Transport{
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return dialer()
		},
	}
	httpClient := &http.Client{Transport: t}
	buf := make([]byte, 1024*1024)
	rand.Read(buf)

	for i := 0; i < 100000; i++ {
		req, err := http.NewRequest("PUT", "https://foo/bar", bytes.NewReader(buf))
		if err != nil {
			panic(err)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Panicf("i: %d, response: %v, err: %v\n", i, resp, err)
		}
	}
}