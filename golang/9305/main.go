package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	postData :=
		`--xxx
Content-Disposition: form-data; name="field1"

value1
--xxx
Content-Disposition: form-data; name="field2"

value2
--xxx
Content-Disposition: form-data; name="file"; filename="file"
Content-Type: application/octet-stream
Content-Transfer-Encoding: binary

binary data
--xxx--
`
	r, _ := http.NewRequest("POST", "/", strings.NewReader(postData))
	r.Header.Set("Content-Type", "multipart/form-data; boundary=xxx")
	err := r.ParseMultipartForm(100000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("    FormValue(field1) =", r.FormValue("field1"))
	fmt.Println("PostFormValue(field1) =", r.PostFormValue("field1"))
}
