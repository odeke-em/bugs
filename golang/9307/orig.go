package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	var mediaURI string
	flag.StringVar(&mediaURI, "uri", "https://sites.ualberta.ca/~odeke/wolves.mp4", "uri of the media to transcode to gif")
	flag.Parse()

	cmd := exec.Command("ffmpeg", "-i", "-", "-f", "gif", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("StdinPipe: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("StdoutPipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatalf("Start: %v", err)
	}
	go func() {
		res, err := http.Get(mediaURI)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode < 200 && res.StatusCode > 299 {
			log.Fatalf("%d %s", res.StatusCode, res.Status)
		}
		log.Printf("fetched\n")

		if _, err := io.Copy(stdin, res.Body); err != nil {
			log.Fatal(err)
		}
		log.Printf("done writing\n")
		if os.Getenv("NO_CLOSING_STDIN") == "" {
			if err := stdin.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}()
	if _, err := io.Copy(os.Stdout, stdout); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Wait: %v", err)
	}
}
