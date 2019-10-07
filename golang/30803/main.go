package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"reflect"
	"time"
)

func h404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(reflect.TypeOf(w))
		w.WriteHeader(404)
		w.WriteHeader(404)
	})
}

func main() {
	cst := httptest.NewServer(http.TimeoutHandler(h404(), time.Second*10, "no logs here"))
	defer cst.Close()

	for i := 0; i < 100; i++ {
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
