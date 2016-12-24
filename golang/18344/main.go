package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{"key1": []string{"a", "b"}, "key2": []string{"1", "2"}}
	fmt.Printf("#%v\n", v)
	v.Del("key1")
	fmt.Printf("#%v\n", v)
}
