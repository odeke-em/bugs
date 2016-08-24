package main

func main() {
    ch := make(chan int)    
    switch {
    case <-ch:
    default:
    }
}
