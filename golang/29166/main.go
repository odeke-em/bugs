package main

import "math"

const (
    a = 0.1
    b = 0.2
    c = 0.3
)

var logA = math.Log(a)
var logC = math.Log(c)

func foo(x float64) float64 {
    return math.Exp(math.Log(a) * math.Log(b * x) / math.Log(c))
}

func bar(x float64) float64 {
    return math.Exp(logA * math.Log(b * x) / logC)
}

func main() {}