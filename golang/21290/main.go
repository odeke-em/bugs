package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "http://ruby.streamguys.com:8120/status2.xsl?mount=/live"

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
}
