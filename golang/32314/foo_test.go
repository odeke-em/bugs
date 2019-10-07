package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConnCloseNoCancellation(t *testing.T) {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		defer func() {
			ctx := r.Context()
			select {
			case <-ctx.Done():
				t.Fatalf("client.Context Done with error: %v", ctx.Err())
			default:
			}
		}()
		if !ok {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		conn.Close()

		reqPayload, _ := ioutil.ReadAll(bufrw)
		t.Logf("client request payload: %s\n", reqPayload)
		bufrw.Write([]byte("HTTP/1.1 200 OK\r\nConnection: keep-alive\r\nContent-Encoding: chunked\r\nContent-Length: 2\r\n\r\nok"))
		bufrw.Flush()
	}))
	defer cst.Close()

	req, _ := http.NewRequest("POST", cst.URL, strings.NewReader("aaaaaaa*****aaaaaaaaaaaa"))
	res, _ := cst.Client().Do(req)
	if res != nil {
		_, _ = ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
	}
}
