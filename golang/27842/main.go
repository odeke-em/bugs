package main

import (
	"strings"

	"golang.org/x/net/html"
)

func main() {
	r := strings.NewReader("<table><math><select><mi><select></table>")
	html.Parse(r)
}
