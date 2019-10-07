package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httptrace"
)

func main() {
	dialFunc := func(ctx context.Context, network string, addr string) (net.Conn, error) {
		// This would be where there's a net.Dial to an external lookup service
		// but just pretend that this actually matters for the purposes of this bug
		// report
		n, err := (&net.Dialer{}).DialContext(ctx, "tcp", "www.golang.org:80")
		if err != nil {
			panic(err)
		}
		n.Close()
		// assume that 1.1.1.1 is an IP returned from the lookup service
		// This IP was purely chosen because it's an IP that I know runs a
		// HTTP server
		return (&net.Dialer{}).DialContext(ctx, network, "1.1.1.1:80")
	}

	transport := &http.Transport{
		DialContext: dialFunc,
	}
	req, err := http.NewRequest("GET", "http://example.org/", nil)
	if err != nil {
		panic(err)
	}
	ct := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Println("DNSStart:", info.Host)
		},
	}
	ctx := httptrace.WithClientTrace(context.Background(), ct)
	req = req.WithContext(ctx)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}
