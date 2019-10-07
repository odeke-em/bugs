// Ensure that a missing status doesn't make the server panic
// See Issue https://golang.org/issues/21701
func TestMissingStatusNoPanic(t *testing.T) {
	t.Parallel()

	ln := newLocalListener(t)
	addr := ln.Addr().String()
	rawRespChan := make(chan string, 1)
	shutdown := make(chan bool, 1)
	done := make(chan bool)
	fullAddrURL := fmt.Sprintf("http://%s", addr)

	go func() {
		defer func() {
			ln.Close()
			close(done)
		}()

		running := true
		for running {
			select {
			case <-shutdown:
				running = false
			case rawResp := <-rawRespChan:
				conn, _ := ln.Accept()
				if conn != nil {
					io.WriteString(conn, rawResp)
					conn.Close()
				}
			}
		}
	}()

	tests := [...]struct {
		raw  string
		want string
	}{
		0: {
			want: "unknown status code",
			raw: `HTTP/1.1 400
Date: Wed, 30 Aug 2017 19:09:27 GMT
Content-Type: text/html; charset=utf-8
Content-Length: 10
Last-Modified: Wed, 30 Aug 2017 19:02:02 GMT
Vary: Accept-Encoding` + "\r\n\r\nAloha Olaa",
		},
		1: {
			want: "malformed HTTP",
			raw: `HTTP/1.1
Date: Wed, 30 Aug 2017 19:09:27 GMT
Content-Type: text/html; charset=utf-8
Content-Length: 10
Last-Modified: Wed, 30 Aug 2017 19:02:02 GMT
Vary: Accept-Encoding` + "\r\n\r\nAloha Olaa",
		},
	}

	proxyURL, err := url.Parse(fullAddrURL)
	if err != nil {
		t.Fatalf("proxyURL: %v", err)
	}

	client := &Client{
		Transport: &Transport{Proxy: ProxyURL(proxyURL)},
	}

	for i, tt := range tests {
		rawRespChan <- tt.raw
		req, _ := NewRequest("POST", "https://golang.org/", nil)
		res, err, panicked := doFetchCheckPanic(client, req)
		if panicked {
			t.Errorf("#%d: panicked, expecting an error", i)
			continue
		}
		if res != nil && res.Body != nil {
			io.Copy(ioutil.Discard, res.Body)
			res.Body.Close()
		}
		if err == nil || !strings.Contains(err.Error(), tt.want) {
			t.Errorf("#%d: got=%v want=%q", i, err, tt.want)
		}
	}
	close(shutdown)
	<-done
}

func doFetchCheckPanic(c *Client, req *Request) (res *Response, err error, panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	res, err = c.Do(req)
	return
}
