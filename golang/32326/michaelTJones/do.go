package main

import (
	"os"
	"runtime"
)

func main() {
        prc, pwc, err := os.Pipe()
        if err != nil {
            panic(err)
        }
        defer pwc.Close()

	ok := make(chan bool)
	go func() {
		bs := make([]byte, 100)
		ok <- true
		_, _ = prc.Read(bs)
	}()

	<-ok

	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			buf = buf[:n]
                        break
		}
		buf = make([]byte, 2*len(buf))
	}

        println(string(buf))
}
