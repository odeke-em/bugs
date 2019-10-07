package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	var workserr, failserr, alsoworkserr error

	input := `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"></s:Envelope>`

	works := struct {
		XMLName xml.Name
	}{}

	fails := struct {
		XMLName xml.Name `xml:"s:Envelope"`
	}{}

	workserr = xml.Unmarshal([]byte(input), &works)
	failserr = xml.Unmarshal([]byte(input), &fails)

	alsoworks, alsoworkserr := xml.Marshal(struct {
		XMLName xml.Name `xml:"s:Envelope"`
	}{})

	fmt.Println("    works:", works.XMLName)
	fmt.Println("    error:", workserr)

	fmt.Println("    fails:", fails.XMLName)
	fmt.Println("    error:", failserr)

	fmt.Println("alsoworks:", string(alsoworks))
	fmt.Println("    error:", alsoworkserr)
}
