package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const html = `
<html>
  Validation
  <form method="POST" action="/validate" enctype="multipart/form-data">
    <input type="file" name="file" />
    <br />
    <input type="submit" value="Send" />
  </form>
</html>
`

func validate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		fmt.Fprintf(w, "parseForm: %v\n", err)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "formFile retrieval %s\n", err)
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%s\n", header.Header)
	fmt.Fprintf(w, "%s\n", header.Filename)
	fmt.Println(header.Header)
	fmt.Println(header.Filename)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html)
}

func main() {
	addr := ":8090"
	http.HandleFunc("/", index)
	http.HandleFunc("/validate", validate)
	if false { // This is your specific command, not present on *NIX
		go exec.Command("rundll32", "url.dll,FileProtocolHandler",
			fmt.Sprintf("http://localhost%s", addr)).Start()
	}
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
