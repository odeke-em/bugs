package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strings"
)

const TEST_BODY = `
Preamble
--MyBoundary
Content-Type: text/plain

This was a triumph.
--MyBoundary

I'm making a note here.
--MyBoundary

HUGE SUCCESS.
--MyBoundary
X-Lie: Cake

It's hard to overstate
my satisfaction.
--MyBoundary--
Postamble

At this point the multipart reader has read all it can.
`

// Returns a reader primed with valid multi-part data.
func NewDummyReader() io.Reader {
	body := strings.Replace(TEST_BODY, "\n", "\r\n", -1)
	out, in := io.Pipe()
	go func() {
		in.Write([]byte(body))
		// Remove the comment below and everything works.
		//in.Close()
	}()
	return out
}

func main() {
	dr := NewDummyReader()
	mr := multipart.NewReader(dr, "MyBoundary")
	for {
		p, err := mr.NextPart()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(p.Header)
		data, err := ioutil.ReadAll(p)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(data))
		fmt.Println()
	}
	fmt.Println("Read all the parts!")
}
