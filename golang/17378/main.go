package main

var x = [1 << 31]byte{0: 1}

func main() {
    println(x[0])
}
