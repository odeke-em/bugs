package main

import (
	"io/ioutil"
	"log"
)

func main() {

	debugf := func(l *log.Logger, format string, v ...interface{}) {
			l.Printf(format, v...)
	}
	l := log.New(ioutil.Discard, "", 0)
	debugf(l, "i=%v", 10)
}

type foo int

func (*foo) Printf(fmt_ string, v ...interface{}) {
}
