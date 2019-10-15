package main

import (
	"io"
        "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"

	"golang.org/x/net/http2"
)

func main() {
	cst := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            blob, _ := httputil.DumpRequest(r, false)
            println(string(blob))
            fmt.Printf("type: %T\n", w)
            w.Header().Set(":protocol", "ws")
            return
		conn, brw, err := w.(http.Hijacker).Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Received request(%s) at: %s\n", conn.RemoteAddr(), conn.LocalAddr())
		blob10, _ := ioutil.ReadAll(io.LimitReader(brw, 10))
		log.Printf("Blob from client: %s\n", blob10)

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
		_ = fr.WriteGoAway(10, http2.ErrCodeNo, []byte("ended here"))
		brw.Flush()
		return
	}))
	if err := http2.ConfigureServer(cst.Config, new(http2.Server)); err != nil {
		log.Fatalf("Failed to configure http2 server: %v", err)
	}
        	cst.TLS = cst.Config.TLSConfig
	cst.StartTLS()
	defer cst.Close()

	tr := &http.Transport{TLSClientConfig: cst.Config.TLSConfig}
	if err := http2.ConfigureTransport(tr); err != nil {
		log.Fatalf("Failed to configure http2 transport: %v", err)
	}
	tr.TLSClientConfig.InsecureSkipVerify = true
	client := &http.Client{Transport: tr}


	req, _ := http.NewRequest("POST", cst.URL, strings.NewReader(`{"ack": true}`))
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return
	}
        blob, _ := httputil.DumpResponse(res, true)
	log.Printf("Response:\n%s\n", blob)
}
