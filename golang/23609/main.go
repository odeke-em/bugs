package main

type Foo struct {
    A string
}

type Bar struct {
    Foo
}

func main() {
    a := Bar{
       A: "bar",
    }
}