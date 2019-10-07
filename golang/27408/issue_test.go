package main_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

func TestIssue(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slurp, _ := ioutil.ReadAll(r.Body)
		_ = r.Body.Close()
		fmt.Fprintf(w, "method: %s, body: %s\n", r.Method, slurp)
	}))
	defer backend.Close()
	bu, _ := url.Parse(backend.URL)

	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}

		bc, err := net.Dial("tcp", bu.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer bc.Close()

		cc, bw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cc.Close()

                if err := r.Write(bc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var wg sync.WaitGroup
		wg.Add(2)
		cp := func(dst io.Writer, src io.Reader) {
			_, _ = io.Copy(dst, src)
			if conn, ok := dst.(interface {
				CloseWrite() error
			}); ok {
				conn.CloseWrite()
			}
			wg.Done()
		}

		lr := io.LimitReader(bw, int64(bw.Reader.Buffered()))
		mr := io.MultiReader(lr, cc)

		go cp(cc, bc)
		go cp(bc, mr)
		wg.Wait()
	}))
	defer proxy.Close()

	// And now for the client

	gotLog := new(bytes.Buffer)
	clientReq := func(msg string) error {
		body := bytes.NewBufferString(msg)
		req, err := http.NewRequest("POST", proxy.URL, body)
		if err != nil {
			return err
		}
		resp, err := proxy.Client().Do(req)
		if err != nil {
			return err
		}
		_, err = io.Copy(gotLog, resp.Body)
		resp.Body.Close()
		return err
	}

	payloads := []string{"Hello", "World"}
	for _, payload := range payloads {
		if err := clientReq(payload); err != nil {
			t.Fatalf("Payload %q ==> %v", payload, err)
		}
	}
        wantLog := "method: POST, body: Hello\nmethod: POST, body: World\n"
	if g, w := gotLog.String(), wantLog; g != w {
		t.Fatalf("\nGot =%q\nWant=%q", g, w)
	}
}
