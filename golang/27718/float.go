package main

//go:noinline
func f(x float32) float32{
	return x + 0
}

func main() {
	var zero float32
	var inf = 1.0 / zero
	var negZero = -1 / inf
	if 1/f(negZero) != inf {
		panic("negZero+0 != posZero")
	}
}
