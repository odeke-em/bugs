package main

func f(p, q *struct{}) bool {
        return *p == *q
}

func main() {
        f(nil, nil)
}