package main

func main() {
	_, _ = Y()
}
func Y() (int, bool) {
	ii := int(1)
	return ii, 0 <= ii && ii < 20
}
