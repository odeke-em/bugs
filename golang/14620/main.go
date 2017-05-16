package main

import "unsafe"

type hdr struct {
	next *msg
}

type msg struct {
	hdr
	pad [1024-hdrsize]uint8
}

const hdrsize = unsafe.Sizeof(hdr{})
