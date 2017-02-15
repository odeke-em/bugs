package main

import (
    "log"
    "net/http"
    "time"
)

func main() {
    srv := &http.Server{
        Addr:        "localhost:8088",
        Handler:     myhandler{},
        ReadTimeout: time.Second,
    }
    go func() {
        // wait server to start
        time.Sleep(time.Second)
        resp, err := http.Get("http://" + srv.Addr)
        if err != nil {
            log.Println(err)
            return
        }
        resp.Body.Close()
    }()
    log.Fatal(srv.ListenAndServe())
}

type myhandler struct{}

func (h myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if cn, ok := w.(http.CloseNotifier); ok {
        <-cn.CloseNotify()
        log.Printf("unexpected closed client")
    }
}
