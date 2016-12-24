package main

func main() {
	_ = []byte("abc", "def", 12)
	_ = string("a", "b")
	_ = []byte()
	_ = string()
	_ = [][]interface{}("abcef", "efg", []string{"a", "b"}, map[string]string{"a": "b"})
	_ = map[string]string([]string{"hey"}, "nil")
}
