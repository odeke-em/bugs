package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	badProxy := nefariousProxy()
	go func() {
		if err := http.ListenAndServe(":8899", badProxy); err != nil {
			log.Fatalf("serving bad proxy err: %v", err)
		}
	}()

	if true {
		rln, err := net.ListenTCP("tcp4", &net.TCPAddr{Port: 9788})
		if err != nil {
			log.Fatal(err)
		}
		ln := &keepAliveListener{rln}
		if http.Serve(ln, http.HandlerFunc(handler)); err != nil {
			log.Fatalf("serve err: %v", err)
		}
	} else {
		ownServer()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("reqIn: %+v\n", r)
	fmt.Fprintf(w, "foo")
}

func ownServer() {
	ln, err := net.Listen("tcp", ":9788")
	if err != nil {
		log.Fatal("ownServer: %v", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("conn err: %v", err)
			continue
		}
		serve(conn)
	}
}

var localURL, _ = url.Parse("http://localhost:9788/")

type nefarious int

func (n nefarious) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.RemoteAddr = ""
	http.Redirect(w, r, "http://localhost:9788", 301)
}

func nefariousProxy() http.Handler {
	return nefarious(0)
	// Strip the host header
	log.Printf("localURL: %v\n", localURL)
	rpxy := httputil.NewSingleHostReverseProxy(localURL)
	if false {
		rpxy.Director = func(r *http.Request) {
			log.Printf("r: %v\n", r)
			delete(r.Header, "Host")
		}
	}
	return rpxy
}

func serve(conn net.Conn) {
	br := bufio.NewReader(conn)
	if false {
		req, err := http.ReadRequest(br)
		if err != nil {
			log.Printf("readRequest err: %v", err)
			return
		} else {
			dump, _ := httputil.DumpRequest(req, true)
			log.Printf("r: %v parsedRequest:\n%s\n", req.RemoteAddr, dump)
		}
	} else {
		blob, _ := ioutil.ReadAll(br)
		log.Printf("blob:\n%s\n", blob)
	}
	fmt.Fprintf(conn, `HTTP/1.1
Date: Wed, 30 Aug 2017 19:09:27 GMT
Content-Type: text/html; charset=utf-8
Content-Length: 10
Last-Modified: Wed, 30 Aug 2017 19:02:02 GMT
Vary: Accept-Encoding`+"\r\n\r\nAloha Olaa",
	)
	conn.Close()
}

type ListenerWrap struct {
	isUnix  bool
	l       net.Listener
	stopped bool
	addrDup string
	addr    string
}

var _ net.Listener = (*ListenerWrap)(nil)

func Listen(proto, addr string) (*ListenerWrap, error) {
	var isUnix bool
	var l net.Listener
	var err error
	var addr_dup string
	if proto == "unix" {
		isUnix = true
		if _, err := os.Stat(addr); !os.IsNotExist(err) {
			os.Remove(addr)
		}
		addr_dup = addr + "_dup"
		if _, err := os.Stat(addr_dup); !os.IsNotExist(err) {
			os.Remove(addr_dup)
		}
	} else if proto == "tcp" {
		// XXX I do not know how to convert ipv4 to ipv6.
		// And net.Listen("tcp", addr) will only listen on ipv6.
		proto = "tcp4"
	}
	if proto == "unix" {
		l, err = net.Listen(proto, addr_dup)
		if err != nil {
			return nil, err
		}
		// Refer to https://github.com/golang/go/issues/11826
		// We don't want Close() delete the file,
		// use hard link to make a duplicated entry,
		// deleting "addr_dup" does not affect "addr"
		err = os.Link(addr_dup, addr)
	} else {
		l, err = net.Listen(proto, addr)
	}
	if err != nil {
		return nil, err
	}
	lw := &ListenerWrap{addr: addr, addrDup: addr_dup, isUnix: isUnix}
	lw.l = l
	return lw, nil
}

func (lw *ListenerWrap) Close() error {
	if lw.stopped {
		return nil
	}
	lw.stopped = true
	return lw.l.Close()
}

func (lw *ListenerWrap) Accept() (c net.Conn, err error) {
	if lw.stopped {
		return
	}
	c, err = lw.l.Accept()
	return
}

func (lw *ListenerWrap) Addr() net.Addr {
	return lw.l.Addr()
}

type keepAliveListener struct {
	*net.TCPListener
}

func (l keepAliveListener) Accept() (_ net.Conn, rerr error) {
	c, err := l.TCPListener.AcceptTCP()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := c.Close(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	f, err := c.File()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	fc, err := net.FileConn(f) //See https://github.com/golang/go/issues/7605 and https://github.com/golang/go/issues/5052
	if err != nil {
		return nil, err
	}
	defer func() {
		if rerr != nil {
			_ = fc.Close() //ignore, rerr already non nil
		}
	}()

	c2 := fc.(*net.TCPConn) // <---- in very rare cases it has everything set correctly but "raddr".

	if err := c2.SetKeepAlive(true); err != nil {
		return nil, err
	}
	if err := c2.SetKeepAlivePeriod(30 * time.Second); err != nil {
		return nil, err
	}

	return c2, nil
}
