package main

import (
        "errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go func() {
			ctx := r.Context()
			select {
			case <-ctx.Done():
				r.Body.Close()
                        default:
			}
		}()
		_, _ = io.Copy(ioutil.Discard, r.Body)
	}))
	defer cst.Close()

	req, err := http.NewRequest("POST", cst.URL, new(neverEndingReader))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	go func() {
		<-time.After(3 * time.Second)
		_ = req.Body.Close()
	}()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	_, _ = io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()
}

type neverEndingReader int

var _ io.Reader = (*neverEndingReader)(nil)

func (ner *neverEndingReader) Read(b []byte) (int, error) {
	copy(b, "aaaaa")
	return 5, nil
}

func (ner *neverEndingReader) Close() error {
	if *ner == 0 {
		*ner = 1
		return io.EOF
	}
	return errors.New("Already closed")
}
