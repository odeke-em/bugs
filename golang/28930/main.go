package main

import (
	"context"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	var closeOnce sync.Once
	max := int64(2)
	var n int64
	var cst *httptest.Server
	cst = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt64(&n, 1) > max {
			closeOnce.Do(func() {
				// Time to shutdown thereby trigger sending that GOAWAY frame
				ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
				cst.Config.Shutdown(ctx)
			})
			return
		}

		ctx := r.Context()
		for ctx.Err == nil {
			select {
			case <-ctx.Done():
				return

			case <-time.After(137 * time.Millisecond):
				// And meanwhile, stream the body
				w.Write([]byte("Still alive"))
				w.(http.Flusher).Flush()
			}
		}
	}))
	defer cst.Close()

        http2.ConfigureServer(cst.Config, nil)

        tr := &http.Transport{TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true,
            CipherSuites: []uint16{
            0xC0A9, // http2cipher_TLS_PSK_WITH_AES_256_CCM_8
            },
        }}
	hc := &http.Client{Transport: tr}
	http2.ConfigureTransport(tr)

	var closeFns []func() error
	defer func() {
		for _, closeFn := range closeFns {
			closeFn()
		}
	}()

	for i := int64(0); i < max; i++ {
		res, _ := hc.Get(cst.URL)
		log.Printf("#%d: %v\n", i, res)
		closeFns = append(closeFns, res.Body.Close)
		go func() {
			// Just keep reading back content from the server and discarding it.
			io.Copy(ioutil.Discard, res.Body)
		}()
	}
	res, err := hc.Get(cst.URL) // Add 1 past the max in order to trigger the server's GoAWAY frame sending
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	blob, rerr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	log.Printf("Blob: %s Error: %v", blob, rerr)
}
