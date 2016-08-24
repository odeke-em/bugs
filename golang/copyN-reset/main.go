package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	start := int(0)
	end := int(0)
	content := ""
	sourceFile := ""
	flag.IntVar(&start, "start", 0, "the offset to start the seek from")
	flag.IntVar(&end, "end", 0, "the offset to start the seek until to")
	flag.StringVar(&content, "content", "still still still got it", "the content to test with")
	flag.StringVar(&sourceFile, "file", "", "the file whose source you want to test with")
	flag.Parse()

	rs := strings.NewReader(content)
	if sourceFile != "" {
		f, err := os.Open(sourceFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	statInfo, err := f.Stat()

	seekStart := int64(start)
	seekEnd := int64(end)

	if seekStart < 0 {
		seekStart = 0
	}

	if seekStart > 0 {
		if _, err := sd.Seek(seekStart, os.SEEK_SET); err != nil {
			log.Fatalf("seek start err=%v, startValue=%d", err, seekStart)
		}
	}

	fn := io.Copy
	if seekEnd > 0 && seekEnd > seekStart {
		diff := seekEnd - seekStart
		fn = func(w io.Writer, r io.Reader) (int64, error) {
			return io.CopyN(w, r, diff)
		}
	}

	fn(os.Stdout, sd)
}
