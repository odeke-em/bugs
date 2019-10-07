package main

import (
	"context"
	"fmt"
	"net"
	"os"
)

func main() {
	hostname := "www.google.com"

	var r *net.Resolver
	//r := &net.Resolver{}
	ips, err := r.LookupIPAddr(context.Background(), hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LookupIPAddr(%q) failed: %s\n", hostname, err)
		os.Exit(1)
	}
	for _, a := range ips {
		fmt.Printf("LookupIPAddr(%q): %s\n", hostname, &a)
	}

	oneIP, err := net.ResolveIPAddr("ip4", hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ResolveIPAddr(%q) failed: %s\n", hostname, err)
		os.Exit(1)
	}
	fmt.Printf("ResolveIPAddr(%q): %s\n", hostname, oneIP)
}
