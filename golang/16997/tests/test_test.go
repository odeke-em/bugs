package err_test

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestDialTransportPreservesNetOpError(t *testing.T) {
	// We need a server, so that we can use its host
	// as a prefix to create a non existent proxy URL.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

	parsedCSTURL, _ := url.Parse(ts.URL)

	segments := strings.Split(parsedCSTURL.Host, ":")
	nonExistentProxyURL := &url.URL{
		Scheme: parsedCSTURL.Scheme,
		Path:   fmt.Sprintf("%s-testing:%s", segments[0], strings.Join(segments[1:], ":")),
	}

	tr := &http.Transport{
		Proxy: func(*http.Request) (*url.URL, error) {
			return url.Parse(nonExistentProxyURL.String())
		},
		Dial: (&net.Dialer{Timeout: time.Nanosecond}).Dial,
	}
	c := http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", "http://golang.org", nil)
	_, err := c.Do(req)
	if err == nil {
		t.Fatal("wanted a non-nil error")
	}

	uerr, ok := err.(*url.Error)
	if !ok {
		t.Fatalf("got %T, want url.Error", err)
	}

	if _, ok := uerr.Err.(*net.OpError); !ok {
		t.Errorf("got %T want net.OpError", uerr)
	}
}
