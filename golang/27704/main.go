package main

import (
	"strings"

	"golang.org/x/net/html"
)

func main() {
	r := strings.NewReader("<template><tBody><isindex/action=0></template>")
	html.Parse(r)
}
