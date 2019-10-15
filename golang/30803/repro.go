package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"reflect"
	"time"
)

        var maxTime = 1 * time.Second
func h404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(reflect.TypeOf(w))
		w.WriteHeader(404)
		w.WriteHeader(404)
		w.WriteHeader(404)
		w.WriteHeader(404)
                <-time.After(maxTime * 2)
	})
}

func main() {
	cst := httptest.NewUnstartedServer(http.TimeoutHandler(h404(), maxTime, "no logs here"))
	cst.Config.ErrorLog = log.New(os.Stdout, "", 0)
	cst.Start()
	defer cst.Close()

	for i := 0; i < 2; i++ {
		res, err := cst.Client().Get(cst.URL)
		if err != nil {
			log.Printf("%d: error: %v", i, err)
			continue
		}

		blob, _ := httputil.DumpResponse(res, true)
		res.Body.Close()
		log.Printf("#%d:\n%s\n\n", i, blob)
	}
}
