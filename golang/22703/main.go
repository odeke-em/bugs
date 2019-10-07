package main
type I interface {
	foo()
}
func main() {
	var i I
	i.foo()
}