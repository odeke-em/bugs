package main

func main() {
	var s []int
	var is interface{} = s
	go func() {
		_ = s == nil
	}()
	go func() {
		_ = is == is
	}()
	
	var m map[string]bool = nil
	_ = m["Go"]
	_, _ = m["Go"]
	m, _ = is.(map[string]bool)
	m = is.(map[string]bool)
}