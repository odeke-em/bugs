package proxy_test

import (
  "net/http"
  "net/http/httptest"
  "net/http/httputil"
  "net/url"
  "testing"
)

func copyHeader(dst, src http.Header) {
  for k, vv := range src {
    for _, v := range vv {
      dst.Add(k, v)
    }
  }
}

func TestHeaders(t *testing.T) {
  // An origin, which should receive a request modified by our reverse proxy director.
  srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if got, want := r.Header.Get("X-Want-Dropped"), ""; got != want {
      t.Errorf("%s: origin request header: X-Want-Dropped: got: %q want: %q", r.URL.Path, got, want)
    }
  }))
  defer srv.Close()
  parsed, err := url.Parse(srv.URL)
  if err != nil {
    t.Fatalf("url.Parse(%q): %v", srv.URL, err)
  }
  // The reverse proxy.
  rp := httputil.NewSingleHostReverseProxy(parsed)
  origDirector := rp.Director
  rp.Director = func(r *http.Request) {
    origDirector(r)
    // Copy header, since r.Header is a shallow copy of the original request
    // headers, and  calls rp.Director for an http.Handler, and http.Handlers
    // aren't permitted to modify the original request.
    h := make(http.Header)
    copyHeader(h, r.Header)
    r.Header = h
    r.Header.Del("X-Want-Dropped")
  }
  rpSrv := httptest.NewServer(rp)
  defer rpSrv.Close()

  tests := []struct {
    label, path      string
    connectionHeader string
  }{
    {"header is dropped with empty connection header", "/without_connection_header", ""},
    {"header is dropped with Connection: close", "/with_connection_header", "close"}, // bug repro.
  }
  for _, test := range tests {
    func() {
      req, err := http.NewRequest(http.MethodGet, rpSrv.URL+test.path, nil)
      if err != nil {
        t.Errorf("%s: http.NewRequest: %v", test.label, err)
        return
      }
      req.Header.Set("X-Want-Dropped", "1")
      if test.connectionHeader != "" {
        req.Header.Set("Connection", test.connectionHeader)
      }
      resp, err := http.DefaultTransport.RoundTrip(req)
      if err != nil {
        t.Errorf("%s: RoundTrip: %v", test.label, err)
        return
      }
      resp.Body.Close()
    }()
  }
}