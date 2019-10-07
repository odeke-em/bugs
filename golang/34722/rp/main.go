package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"

	"golang.org/x/net/http2"
)

func main() {
	cst := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		files := []string{"main.go"}
		if err := serveDownload(r.Context(), w, files); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	if err := http2.ConfigureServer(cst.Config, new(http2.Server)); err != nil {
		log.Fatal(err)
	}
	cst.TLS = cst.Config.TLSConfig
	cst.StartTLS()
	defer cst.Close()

	tr := &http.Transport{
		TLSClientConfig: cst.Config.TLSConfig,
	}
	tr.TLSClientConfig.InsecureSkipVerify = true
	http2.ConfigureTransport(tr)
	req, _ := http.NewRequest("GET", cst.URL, nil)

	res, err := tr.RoundTrip(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer res.Body.Close()

	if false { // Uncomment in order to examine the entire response and ensure HTTP/2.0 being used.
		blob, _ := httputil.DumpResponse(res, true)
		println(string(blob))
	} else {
		io.Copy(os.Stdout, res.Body)
	}
}

func serveDownload(ctx context.Context, w http.ResponseWriter, files []string) (err error) {
	// Defines the file name.
	name := "archive.zip"
	w.Header().Set("Content-Disposition", "attachment; filename*=utf-8''"+url.PathEscape(name))
	pr, pw := io.Pipe()
	//j - cut path, and store only
	cmd := exec.Command("zip", append([]string{"-0rq", "-"}, files...)...)
	cmd.Stdout = pw
	cmd.Stderr = os.Stderr
	go func() {
		defer pr.Close()
		// copy the data written to the PipeReader via the cmd to stdout
		if _, err := io.Copy(w, pr); err != nil {
			panic(err)
		}
	}()
	return cmd.Run()
}
