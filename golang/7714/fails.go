package main
func main() {
    var y int
    set := func(p *int, x int) {
        *p = x
    }
    set(&y, 42)
}
