package main

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"

	"golang.org/x/net/http2"
)

func main() {
	cst := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, brw, err := w.(http.Hijacker).Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tlc := conn.(*net.TCPConn)
		tlc.SetLinger(0)
		log.Printf("Received request(%s) at: %s\n", tlc.RemoteAddr(), tlc.LocalAddr())
		blob10, _ := ioutil.ReadAll(io.LimitReader(brw, 10))
		log.Printf("Blob from client: %s\n", blob10)

		if false {
			brw.Write([]byte("HTTP/2.0 200 OK\r\nConnection: keep-alive\r\nContent-Encoding: chunked\r\nContent-Length: 2\r\n\r\nok"))
			brw.Flush()
			return
		}

		fr := http2.NewFramer(brw, brw)
		_ = fr.WritePing(true, [8]byte{'h', 'e', 'l', 'l', 'o', 'o', 'l', 'a'})
		unknownFrame := http2.FrameContinuation * 2
		flag := http2.FlagDataPadded
		streamID := uint32(1)
		payload := []byte("Hello world. This is an unknown frame")
		if err := fr.WriteRawFrame(unknownFrame, flag, streamID, payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		brw.Flush()
		return
		_ = fr.WriteGoAway(10, http2.ErrCodeNo, []byte("ended here"))
	}))
	if err := http2.ConfigureServer(cst.Config, nil); err != nil {
		log.Fatalf("Failed to configure http2 server: %v", err)
	}
	cst.Start()
	defer cst.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			NextProtos:         []string{"h2"},
			InsecureSkipVerify: true,
		},
	}
	if err := http2.ConfigureTransport(tr); err != nil {
		log.Printf("Failed to configure the transport: %v", err)
		return
	}

	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", cst.URL, strings.NewReader(`{"ack": true}`))
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return
	}
	blob, _ := ioutil.ReadAll(res.Body)
	log.Printf("Response:\n%s\n", blob)
}
