package main

import "runtime"

type node struct {
        next *node
}

var x bool

func main() {
        var head *node
        for x {
                head = &node{head}
        }

        runtime.KeepAlive(head)
}