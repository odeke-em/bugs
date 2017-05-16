package main

func main() {
    ch := make(chan string)
    switch {
    case <-ch:
    default:
    }
}
