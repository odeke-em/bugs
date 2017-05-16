package main

import (
	"net/http"
	"bufio"
	"bytes"
	"os"
)

func main() {
	rspbytes := []byte("HTTP/1.0 200 OK\r\nConnection: keep-alive\r\nContent-Length: 4\r\n\r\nAAAA")
	buf := bytes.NewBuffer(rspbytes)
	
	httpRsp, _ := http.ReadResponse(bufio.NewReader(buf), nil)
	buf2 := bytes.NewBuffer(make([]byte, 0))
	httpRsp.Write(buf2)
	
	httpRsp2, _ := http.ReadResponse(bufio.NewReader(buf2), nil)
	buf3 := bytes.NewBuffer(make([]byte, 0))
	httpRsp2.Write(buf3)
	
	httpRsp3, _ := http.ReadResponse(bufio.NewReader(buf3), nil)
	httpRsp3.Write(os.Stdout)
}
