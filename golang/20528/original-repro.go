package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
)

func handle(resp http.ResponseWriter, req *http.Request) {
	data := map[string]string{}
	for i := 0; i < 600; i++ {
		k := fmt.Sprintf("key%d", i)
		v := fmt.Sprintf("value%d", i)
		data[k] = v
	}

	err := json.NewEncoder(resp).Encode(&data)
	if err != nil {
		panic(err)
	}
}

type mylistener struct {
	net.Listener
	count int
}

func (l *mylistener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	l.count++
	if l.count > 1 {
		panic("more than one connection")
	}
	println("openning connection")
	return c, err
}

func main() {
	ln, err := net.Listen("tcp4", ":0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	srv := httptest.NewUnstartedServer(http.HandlerFunc(handle))
	srv.Listener = &mylistener{ln, 0}
	srv.StartTLS()

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	addr := srv.Listener.Addr()
	for i := 0; i < 2; i++ {
		requestData(client, "https://"+addr.String())
	}
}

func requestData(client http.Client, url string) {
	println("requesting data")
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	m := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		panic(err)
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
}
