package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	examples := []string{
		"http://www.xmlfiles.com/examples/note.xml",
		// "https://www.gpo.gov/fdsys/bulkdata/BILLS/115/1/hconres/BILLS-115hconres18rds.xml",
		// "https://www.w3schools.com/xml/cd_catalog.xml",
	}
tr := &http.Transport{
DisableCompression: true,
}
http.DefaultClient.Transport = tr
	for i := range examples {
		resp, err := http.Get(examples[i])
		if err != nil {
			log.Fatalln(err)
		}
		buf, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("#%d: %s\n", i, buf)
		resp.Body.Close()
	}

	for i := range examples {
		resp, err := http.Head(examples[i])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp.ContentLength)
		resp.Body.Close()
	}
}
