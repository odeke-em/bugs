package main

import (
	"time"
	"net/http"
	"log"
	"strconv"
	"runtime/debug"
)

func main() {
	url := []string{
		"www.google.com",
		"www.abcabc.com",
		"www.github.com",
		"www.randomwebsite.com",
	}
	go func() {
		for {
			<-time.After(time.Second * 1)
			for key := range url {
				go func(key int) {
					startDate := time.Now().Local().String()
					client := http.Client{Timeout: time.Duration(time.Second * 10)}
					request, err1 := http.NewRequest("GET", "http://" + url[key], nil)
					response, err2 := client.Do(request)
					finishDate := time.Now().Local().String()
					if err1 != nil || err2 != nil {
						log.Println("This job for " + url[key] + " has error1:", err1)
						log.Println("This job for " + url[key] + " has error2:", err2)
					} else {
						log.Println("This job for " + url[key] + " Result is " + strconv.Itoa(response.StatusCode) + " from " + startDate + " to " + finishDate)
					}
					if response != nil {
						response.Body.Close()
					}
				}(key)
			}
			debug.FreeOSMemory()
			//runtime.GC()
		}
	}()
	select {}
}