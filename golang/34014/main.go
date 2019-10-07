package main


func main() {
    f2()
}

func f1(x *[]byte, y *[]byte) {
    f2()
}

func f2() {
    var x, y []byte
    foo(&x, &y)
    foo(&x, &y)
}

func foo(x *[]byte, y *[]byte) {
    *x = *y
}
