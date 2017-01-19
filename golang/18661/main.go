package main

func test(obj interface{}) {
        if obj != struct{a *string}{} {
        }
}

func main() {}