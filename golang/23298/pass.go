package p

type T string

func foo() {
	var _ = []rune(T("T"))
}
