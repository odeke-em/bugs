package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	conn, err := tls.Dial("tcp", "gfs.line.naver.jp:443", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatalf("Failed to dial to server: %v", err)
	}
	defer conn.Close()

	// Now write the request on the wire.
        reqPayload := fmt.Sprintf("POST /enc HTTP/1.1\r\nHost: gfs.line.naver.jp\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36\r\naccept: application/x-thrift; charset=latin1\r\ncontent-type: application/x-thrift\r\n"+
"x-lal: en\r\n"+
"x-lcs: 0005fF9u99g+fCJaX0CiyHrj+/nhwDV2k72si+qxjA9ZNcR/zlgE4W1i1DQL3/zhNBC4UDFlwGz5K+TN9iB0L3AhBJ8+FA9OrY9P/4CHtkrqvJmubO37mXmlXnNCyYk2vbb864it3sxbcl5B1HahIh9SkfHgZUvKek6hcobn/c3+GG033xuYvjBt/DzgfrkPcj4Gu3SFxKSGppTobvtkT/YZocTz+U9Hm0DalRsQr1DadjreZgUdSBlAv2DLd59SqEbsx8fIXIvkHZ8s/84ujTEIeA6t9mv2YA0GFobqejP0PSlNvlW5rN+EN/XVdfbF9YKyNxJIzjjFkPpy/F/1D3T13A==\r\n"+
"x-le: 18\r\n"+
"x-lhm: POST\r\n"+
"x-line-application: ANDROIDLITE	8.2.4	Android	OS	6.0\r\n"+
"x-lst: 1000\r\n\r\n%s", payload)
    fmt.Printf("Request dump:\n%q\n", reqPayload)

        nw, err := conn.Write([]byte(reqPayload))
        fmt.Printf("Bytes sent: %d err: %v\n", nw, err)

	// Now read back the response.
        nr, err := io.Copy(os.Stdout, conn)
	log.Printf("Bytes read: %d err: %v\n", nr, err)
}

var payload = []byte{106, 204, 139, 83, 126, 106, 65, 96, 38, 202, 93, 211, 200, 152, 32, 20, 244, 204, 205, 230, 148, 131, 53, 220, 206, 144, 100, 11, 173, 82, 161, 230, 225, 166, 91, 215, 34, 78, 211, 88, 149, 111, 42, 6, 176, 239, 98, 100, 77, 158, 167, 94, 114, 197, 173, 160, 21, 2, 152, 212, 186, 134, 35, 166, 20, 95, 128, 59, 230, 149, 82, 59, 234, 80, 0, 241, 165, 53, 90, 231, 34, 163, 145, 90, 175, 184, 201, 133, 21, 138, 158, 73, 244, 181, 48, 49, 20, 237, 54, 122, 159, 94, 5, 62, 239, 221, 100, 234, 110, 106, 250, 188, 80, 209, 119, 255, 68, 227, 207, 17, 84, 13, 210, 1, 37, 67, 2, 13, 65, 1, 44, 155, 218, 247, 157, 190, 168, 111, 89, 182, 40, 105, 189, 90, 193, 67, 251, 218}
