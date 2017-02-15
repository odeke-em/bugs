package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := httputil.DumpRequest(r, false)
		fmt.Println(string(data))
	}))
	c := http.Cookie{
		Name:  "name1",
		Value: `abc"withquotesinner"outerpair`,
		Path:  "/",
	}
	r, _ := http.NewRequest("GET", s.URL, nil)
	r.AddCookie(&c)
	r.AddCookie(&http.Cookie{
		Name:  "name2",
		Value: `quotednewonethere`,
		Path:  "/",
	})

	r.Header.Add("Cookie", `name2="quoted"influence"hey`)
	_, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
}
