package main

type P interface {
	p() interface {
		PQ
	}
}
type Q interface {
	q() interface {
		PQ
	}
}
type PQ interface {
	P
	Q
}

var pq PQ