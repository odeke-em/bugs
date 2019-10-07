package p
type T string
func foo() {
	var t = T("T")
	_ = []rune(t)
}
