package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("echo${IFS}hello${IFS}$USER;echo${IFS}https://>/dev/null")
	fmt.Println(u, err)
}
