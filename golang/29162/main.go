package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"strings"
	"time"
)

func main() {
	uri := "https://www.azulseguros.com.br/azul-cms/wp-content/themes/azul2016/integracao-azul-leve.php?timestamp=NGJYHX7MXNS"
	req, err := http.NewRequest("GET", uri, nil)
	netClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: &transferEncodingRedactor{base: http.DefaultTransport.(*http.Transport)},
	}
	netClient.Transport = nil
	res, err := netClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	blob, _ := httputil.DumpResponse(res, true)
	fmt.Printf("%s\n", blob)
	_ = res.Body.Close()
}

type transferEncodingRedactor struct {
	base *http.Transport
}

func (rt *transferEncodingRedactor) RoundTrip(req *http.Request) (*http.Response, error) {
	// Optimistic route i.e common case.
	res, err := rt.base.RoundTrip(req)
	if err == nil || res != nil {
		return res, err
	}

	// An error occured here but now time ot examine the contents
	if !strings.Contains(err.Error(), "unsupported transfer encoding") {
		return res, err
	}

	u := req.URL
	var conn net.Conn
	switch req.URL.Scheme {
	case "https":
		tr := rt.base
		if tr == nil {
			tr = http.DefaultTransport.(*http.Transport)
		}
		conn, err = tls.Dial("tcp", "www.azulseguros.com.br:443", tr.TLSClientConfig)
	default:
		conn, err = net.Dial("tcp", "www.azulseguros.com.br:80")
	}

	if err != nil {
		return nil, fmt.Errorf("Dial error: %v", err)
	}

	// Ensure we close the connection after.
	// Perhaps you might want to reuse it per host?
	defer conn.Close()
	br := bufio.NewReader(conn)
	bw := bufio.NewWriter(conn)

	s := u.Path
	if q := u.RawQuery; len(q) > 0 {
		s += "?" + q
	}

	intro := fmt.Sprintf("%s %s HTTP/1.1\r\nHost: %s\r\n\r\n", req.Method, s, u.Host)
	fmt.Fprint(bw, intro)
	bw.Flush()

	// And now we've got the case of an unsupported transfer
	// encoding time for a raw fetch and header rewrite.
	tp := textproto.NewReader(br)

	protoLine, err := tp.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("ReadLine error: %v", err)
	}
	// Time to parse the headers
	rhdr, err := tp.ReadMIMEHeader()
	if err != nil {
		return nil, fmt.Errorf("ReadMIMEHeader: %v", err)
	}
	hdr := http.Header(rhdr)
	// Now rewrite the header
	if te := hdr.Get("Transfer-Encoding"); te == "gzip" {
		hdr.Set("Content-Encoding", "gzip")
		delete(hdr, "Transfer-Encoding")
	}

	// After this rewrite, reassemble the already read parts back
	// and then get them to ReadResponse
	hdrBuf := new(bytes.Buffer)
	if err := hdr.Write(hdrBuf); err != nil {
		return nil, fmt.Errorf("Rewriting header to wire: %v", err)
	}

	reconstructed := io.MultiReader(strings.NewReader(protoLine+"\r\n"), hdrBuf, strings.NewReader("\r\n"), br)
	return http.ReadResponse(bufio.NewReader(reconstructed), req)
}
