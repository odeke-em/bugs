package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	errsChan := make(chan error, 1)
	cst := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		_, err := rw.Write([]byte("Hello, Worlds!"))
		errsChan <- err
	}))
	cst.Config.WriteTimeout = 1 * time.Second
	cst.Start()
	errLog := new(bytes.Buffer)
	cst.Config.ErrorLog = log.New(errLog, "", 0)
	defer cst.Close()

	res, err := cst.Client().Get(cst.URL)
	if err != nil {
            ne := err.(net.Error)
            t.Fatalf("Failed to invoke server: %v timeout: %t", err, ne.Timeout())
	}

	var blob []byte
	if res != nil {
		blob, _ = ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
	}
	err = <-errsChan
	if err == nil {
		t.Error("Expected a non-nil error")
	}
	t.Logf("blob: %s\nerr: %v", blob, err)
}
