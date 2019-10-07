package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type stringHandler string

func (s stringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Result", string(s))
}

func TestMuxer(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("example.org/pkg/foo/", stringHandler("example.org/pkg/foo/"))

	tests := []struct {
		url      string
		code     int
		response string
	}{
		{"http://example.org/pkg/foo", 301, ""},
		{"http://example.org/pkg/foo/", 200, "example.org/pkg/foo/"},
	}

	for i, tt := range tests {
		r, _ := http.NewRequest("GET", tt.url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, r)

		if got, want := w.Code, tt.code; got != want {
			t.Errorf("#%d Status = %d; want = %d", i, got, want)
		}

		if w.Code == 200 {
			if got, want := w.HeaderMap.Get("Result"), tt.response; got != want {
				t.Errorf("$%d Result = %q; want = %q", i, got, want)
			}
		}
	}
}