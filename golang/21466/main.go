package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"

	"golang.org/x/net/http2"
)

func main() {
	t := &http2.Transport{}
	cconn, sconn := net.Pipe()
	go brokenServer(cconn)
	cc, err := t.NewClientConn(sconn)
	if err != nil {
		log.Fatalf("NewClientConn: %v", err)
	}
	log.Print("client: sending request")
	res, err := cc.RoundTrip(&http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/"},
	})
	if err != nil {
		log.Fatalf("client: RoundTrip error: %v", err)
	}
	log.Print("client: RoundTrip response: %v", res)
}

func brokenServer(conn net.Conn) {
	buf := make([]byte, len(http2.ClientPreface))
	if _, err := io.ReadFull(conn, buf); err != nil {
		log.Fatalf("server: read preface: %v", err)
	}
	log.Printf("server: got preface: %q", buf)
	wb := bufio.NewWriter(conn)
	fr := http2.NewFramer(wb, conn)
	log.Print("server: write SETTINGS frame")
	if err := fr.WriteSettings(
		http2.Setting{http2.SettingMaxConcurrentStreams, 10},
		http2.Setting{http2.SettingMaxHeaderListSize, 16384}); err != nil {
		log.Fatalf("server: WriteSettings: %v", err)
	}
	for {
		log.Print("server: waiting for frame")
		f, err := fr.ReadFrame()
		if err != nil {
			log.Printf("server: ReadFrame: %#v", err)
		}
		log.Printf("server: got frame: %#v", f)
		switch f := f.(type) {
		case *http2.SettingsFrame:
			log.Printf("server: sending SETTINGS ACK frame")
			fr.WriteSettingsAck()
		case *http2.HeadersFrame:
			log.Printf("server: sending DATA frame")
			fr.WriteData(f.StreamID, false, []byte("boom"))
			wb.Flush()
		}
	}
}
