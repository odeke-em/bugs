package main

func main() {
	_ = []byte("abc", "def", 12)
	_ = string("a", "b")
	_ = []byte()
	_ = string()
	_ = [][]interface{}("abcef", "efg", []string{"a", "b"})
}
