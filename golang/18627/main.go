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
		Value: `"quoted"`,
		Path:  "/",
	}
	r, _ := http.NewRequest("GET", s.URL, nil)
	r.AddCookie(&c)
	r.Header.Add("Cookie", `name2="quoted"`)
	_, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
}