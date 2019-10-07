package main
type A struct {}
func (A) F() {}
type B struct {}
func (B) F() {}
type C struct {
    A
    B
}
type G struct { A }
var _ = G.F
var _ = C.F
func main() {}
