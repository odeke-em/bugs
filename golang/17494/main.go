package main

import (
	"fmt"
	"time"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	// Start a trivial server.
	go func() {
		http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
			fmt.Println("Got cookie:", req.Header.Get("Cookie"))

			// Set the cookie to a new value.
			http.SetCookie(resp, &http.Cookie{
				Name:  "YumYumCookie",
				Value: "NewValue",
				Path:  "/",
			})

			// Keep redirecting to yourself until you see the new cookie value.
			ck, _ := req.Cookie("YumYumCookie")
			if ck.Value != "NewValue" {
				http.Redirect(resp, req, "/", http.StatusFound)
			}
		})
		http.ListenAndServe("localhost:8888", nil)
	}()

	time.Sleep(time.Second)

	// Make a request to the server. Initialize the request with an old value for the cookie.
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	req, _ := http.NewRequest("GET", "http://localhost:8888/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "YumYumCookie",
		Value: "OldValue",
		Path:  "/",
	})
	req.Header.Add("HeaderKey", "HeaderValue")
	fmt.Println(client.Do(req))
}
