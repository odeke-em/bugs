// TestTransportTimeoutServerHangs demonstrates that client can hang forever
// when the server hangs and the headers exceed the conn buffer size (forcing a
// flush). Without respect to the context's deadline.
func TestTransportTimeoutServerHangs(t *testing.T) {
	clientDone := make(chan struct{})
	ct := newClientTester(t)
	ct.client = func() error {
		defer ct.cc.(*net.TCPConn).CloseWrite()
		defer close(clientDone)

		buf := make([]byte, 1<<19)
		_, err := rand.Read(buf)
		if err != nil {
			t.Fatal("fail to gen random data")
		}
		headerVal := hex.EncodeToString(buf)

		req, err := http.NewRequest("PUT", "https://dummy.tld/", nil)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		req = req.WithContext(ctx)
		req.Header.Add("Authorization", headerVal)
		_, err = ct.tr.RoundTrip(req)
		if err == nil {
			return errors.New("error should not be nil")
		}
		if ne, ok := err.(net.Error); !ok || !ne.Timeout() {
			return fmt.Errorf("error should be a net error timeout was: +%v", err)
		}
		return nil
	}
	ct.server = func() error {
		ct.greet()
		select {
		case <-time.After(5 * time.Second):
		case <-clientDone:
		}
		return nil
	}
	ct.run()
}