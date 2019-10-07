package main

import "time"

func main() {
	c := make(chan time.Time)
	stop := make(chan bool, 1)
	go func() {
		defer close(stop)
		t := &time.Ticker{C: c}
		t.Stop()
	}()
	<-stop
}
