package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func envOrDefault(envVar, fallback string) string {
	if retr := os.Getenv(envVar); retr != "" {
		return retr
	}
	return fallback
}

func main() {
	var dto = &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{"foo", "bar"}

	data, err := json.Marshal(dto)
	log.Printf("err: %v\n", err)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	url := envOrDefault("OUT_URL", "https://api.pipedrive.com/v1/authorizations")
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		// fails here
		log.Fatal(err)
	}

	r, _ := httputil.DumpResponse(res, true)
	log.Println(string(r))
}
