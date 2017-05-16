package main

import (
	"fmt"
	"time"
)

type Result struct {
	Code      uint16        `json:"code"`
	Timestamp time.Time     `json:"timestamp"`
	Latency   time.Duration `json:"latency"`
	BytesOut  uint64        `json:"bytes_out"`
	BytesIn   uint64        `json:"bytes_in"`
	Error     string        `json:"error"`
}

type Results []Result

func (rs Results) Len() int { return len(rs) }

func main() {
	rs := []Result{
		{Code: 10},
		{BytesOut: 15},
	}

	rsl := Results(rs)
rslPtr := &rsl
	fmt.Printf("len: %d\n", rslPtr.Len())
}
