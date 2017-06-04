package main

func main() {
	_ = int(5)               // ok
	_ = struct{ int }{}      // ok
	_ = struct{ int }{5}     // ok
	var _ = int(5)           // ok
	var _ = struct{ int }{}  // ok
	var _ = struct{ int }{5} // cannot use _ as value
}
