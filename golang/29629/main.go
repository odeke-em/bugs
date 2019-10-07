package main

type stack []int

func (s *stack) last() int {
    return (*s)[len(*s)-1]
}

func main() {
    s := stack{}
    s.last() // should panics
}
