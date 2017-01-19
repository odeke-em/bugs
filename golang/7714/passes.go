package main

func set(p *int, x int) {
    *p = x
}

func main() {
    var y int
    set(&y, 42)
}
