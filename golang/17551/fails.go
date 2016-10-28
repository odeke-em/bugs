package main

func main() {
	_, _ = Y()
}
func Y() (i int, ok bool) {
	ii := int(1)
	return ii, 0 <= ii && ii < 20
}
