package main

import "../b"

func main() {
type st struct{ x, y int }
	_ = b.GetY().(st)
