package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anotherGoogleFan/log"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/video/upload", http.HandlerFunc(UploadVideo))
	srv := &http.Server{
		Addr:    ":8848",
		Handler: mux,
	}
	// 如果接收到关闭信号(如 kill -2)则进行graceful shutdown
	go gracefulShutdown(srv)

	// service connections
	if err := srv.ListenAndServe(); err != nil {
		log.Fields{
			"port": "8848",
			"err":  err.Error(),
		}.Error("unexpected service stop")
		return
	}
}

func gracefulShutdown(srv *http.Server) {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, os.Kill)
	<-stopChan
	log.Info("receive shutdown signal. Ready to exist")
	defer log.Info("program exists")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
}

func UploadVideo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "ok", http.StatusOK)
}