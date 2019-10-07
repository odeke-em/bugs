package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/http2"
)

var paths []string

func pushHandler(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		blob, _ := httputil.DumpRequest(request, true)
		fmt.Printf("\nRequest In:\n%s\n", blob)
	}()

	pusher, supported := writer.(http.Pusher)
	if !supported {
		writer.WriteHeader(http.StatusHTTPVersionNotSupported)
		writer.Write([]byte("HTTP/2 push not supported."))
		return
	}

	html := "<html><body><h1>Golang Server Push</h1>"

	errCount := 0
	for _, path := range paths {
		log.Printf("Pushing %s...\n", path)
		if err := pusher.Push(path, nil); err != nil {
			errCount++
			log.Printf("Failed to push: %v", err)
			continue
		}
		html += fmt.Sprintf("<video src=\"%s\"></video>", path)
	}

	fmt.Fprint(writer, html)
}

type responseWriterLogger struct {
	http.ResponseWriter
	status int
}

func (w *responseWriterLogger) WriteHeader(status int) {
	log.Printf("Writing status %d\n", status)
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func wrapFileServer(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wl := &responseWriterLogger{ResponseWriter: w}
		log.Printf("[...] %s %s\n", r.Method, r.URL)
		next.ServeHTTP(wl, r)
		log.Printf("[%d]: %s %s\n", wl.status, r.Method, r.URL)
	}
}

func main() {
	// 1. Generate the files.
	n := 30
	publicDir := "public"
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		log.Fatalf("Failed to create the the directory for the files: %v", err)
	}
	// Remember to clean up on exit.
	defer os.RemoveAll(publicDir)

	for i := 0; i < n; i++ {
		fullPath := filepath.Join(publicDir, fmt.Sprintf("%d.txt", i))
		payload := strings.Repeat(fmt.Sprintf("%d", i+1), 300<<10) // Make them 300kB
		if err := ioutil.WriteFile(fullPath, []byte(payload), 0655); err != nil {
			log.Fatalf("Failed to create %s: %v", fullPath, err)
		}
		paths = append(paths, filepath.Join("/", fullPath))
	}

	// 2. Now run the server.
	relPublicDir := "/" + publicDir + "/"
	mux := http.NewServeMux()
	mux.Handle(relPublicDir, wrapFileServer(http.StripPrefix(relPublicDir, http.FileServer(http.Dir(publicDir+"/")))))
	mux.HandleFunc("/push", pushHandler)

	cst := httptest.NewUnstartedServer(mux)
	http2.ConfigureServer(cst.Config, new(http2.Server))
	cst.TLS = cst.Config.TLSConfig
	cst.StartTLS()
	defer cst.Close()

	ux, _ := url.Parse(cst.URL)
	// 3. Initiate a request using the HTTP/2.0 Go client.
	conn, err := tls.Dial("tcp", ux.Host, &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"h2"},
	})
	if err != nil {
		log.Fatalf("Failed to dial to the server: %v", err)
	}
	defer conn.Close()
	h2Conn, err := new(http2.Transport).NewClientConn(conn)
	if err != nil {
		log.Fatalf("Failed to create http2 clientConn: %v", err)
	}
	defer h2Conn.Close()

	if false {
		fr := http2.NewFramer(conn, conn)
		fr.WriteSettings(http2.Setting{
			ID:  http2.SettingEnablePush,
			Val: 0x2,
		})
	}

	ctx := context.Background()
	if err := h2Conn.Ping(ctx); err != nil {
		log.Fatalf("Failed to send Ping frame: %v", err)
	}

	req, _ := http.NewRequest("GET", cst.URL+"/push", nil)
	res, err := h2Conn.RoundTrip(req)
	if err != nil {
		log.Fatalf("Failed to make client HTTP request: %v", err)
	}
	blob, _ := httputil.DumpResponse(res, true)
	fmt.Printf("blob: %s\n", blob)
	res.Body.Close()
	fmt.Printf("\n\nAnd the server URL is: %s\n", cst.URL)

	// Remove to end repro immediately or keep to test the server
	// out with your browser or HTTP clients.
	for {
		<-time.After(10 * time.Second)
	}
}
