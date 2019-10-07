package deadreq

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"unsafe"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func BenchmarkDeadreq(b *testing.B) {
	client := new(http.Client)
	b.RunParallel(func(pb *testing.PB) {
		buf := make([]byte, 16<<10)
		for pb.Next() {
			req, err := http.NewRequest("GET", "http://www.example.com", nil)
			chk(err)
			fmt.Println("my thinger", unsafe.Pointer(req))
			resp, err := client.Do(req)
			chk(err)
			_, err = io.CopyBuffer(ioutil.Discard, resp.Body, buf)
			chk(resp.Body.Close())
			chk(err)
		}
	})
}
