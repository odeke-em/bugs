package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func exitIfErr(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(-1)
}

func main() {
	if len(os.Args) < 2 {
		exitIfErr(fmt.Errorf("expecting atleast one arg"))
	}

	rest := os.Args[1:]
	for _, filename := range rest {
		f, err := os.Open(filename)
		exitIfErr(err)

		fields := map[string]string{
			"filename": filename,
		}
		res, err := multipartUpload("http://localhost:8090/validate", f, fields)
		_ = f.Close()
		exitIfErr(err)

		io.Copy(os.Stdout, res.Body)
		_ = res.Body.Close()
	}
}

func createFormFile(mw *multipart.Writer, filename string) (io.Writer, error) {
	return mw.CreateFormFile("file", filename)
}

func multipartUpload(destURL string, f io.Reader, fields map[string]string) (*http.Response, error) {
	if f == nil {
		return nil, fmt.Errorf("bodySource cannot be nil")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := createFormFile(writer, fields["filename"])
	if err != nil {
		return nil, fmt.Errorf("createFormFile %v", err)
	}

	n, err := io.Copy(fw, f)
	if err != nil && n < 1 {
		return nil, fmt.Errorf("copying fileWriter %v", err)
	}

	for k, v := range fields {
		_ = writer.WriteField(k, v)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("writerClose: %v", err)
	}

	req, err := http.NewRequest("POST", destURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	if req.Close && req.Body != nil {
		defer req.Body.Close()
	}

	return http.DefaultClient.Do(req)
}
