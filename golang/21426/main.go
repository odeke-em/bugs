package main

import (
	"net/http"
	"net/http/httptest"
	"fmt"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "http://www.qiniu.com/")
		w.WriteHeader(http.StatusSeeOther)
	}))
	defer ts.Close()

	resp, err := http.Post(ts.URL,"multipart/form-data", nil)
	if err != nil {
		fmt.Printf("Post fail, ", err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Fail follow 301, https\n")
	} else {
		fmt.Println("OK")
	}
}