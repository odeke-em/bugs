package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func echoTime(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, fmt.Sprintf("%s", time.Now()))
}

func ensurePortAsAddr(addr string) string {
	return fmt.Sprintf(":%s", strings.TrimLeft(strings.TrimRight(addr, ":"), ":"))
}

func main() {
	portStr := ""
	flag.StringVar(&portStr, "port-addr", ":6789", "port to run server on")
	flag.Parse()

	srv := new(http.Server)
	srv.Addr = ensurePortAsAddr(portStr)

	log.Printf("serving at %s\n", srv.Addr)
	http.HandleFunc("/", echoTime)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
