package main

import "net"

func main() {
	d := &net.Dialer{}
	deadline := d.deadline()
}
