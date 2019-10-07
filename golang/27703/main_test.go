package main

import (
	"log"
	"testing"
	"io/ioutil"
)

func BenchmarkDebugf(b *testing.B) {
	debugf := func(l *log.Logger, format string, v ...interface{}) {
            if l != nil {
			l.Printf(format, v...)
                    }
	}

	bench := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			debugf("i=%v", i)
		}
	}

	l := log.New(ioutil.Discard, "", 0)
	b.Run("No Logger", bench)
	l = log.New(ioutil.Discard, "", 0)
	b.Run("Logger", bench)
}

func main() {
	var l *log.Logger

	debugf := func(format string, v ...interface{}) {
		if l != nil {
			l.Printf(format, v...)
		}
	}
        debugf("i=%v", 10)
}
