package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type content struct {
	TplEncoded      string
	ManuallyEncoded template.URL

	ShowPaths bool
	RawPath   string
	Path      string
}

func main() {
	tpl, _ := template.New("test").Parse(`<!doctype html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<meta charset="utf-8" />
	</head>
	<body>
		{{ if .ShowPaths }}
			<p>RawPath = {{ .RawPath }}</p>
			<p>Path = {{ .Path }}</p>
		{{ else }}
			<a href="/link/{{ .TplEncoded }}">Template encoded link</a><br />
			<a href="/link/{{ .ManuallyEncoded }}">Manually encoded link</a>
			<br />
			<p>View this page's source to see the (lower/upper)case difference
			in the links</p>
		{{ end }}
	</body>
	</html>`)

	// Renders the root with good and bad links.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := "😋" // Unicode emoji.
		tpl.Execute(w, content{
			// html/template encodes into lowercase characters.
			TplEncoded: s,

			// url.PathEscape encodes into uppercase characters.
			ManuallyEncoded: template.URL(url.PathEscape(s)),
		})
	})

	// This handler produces inconsistent RawPath based on (upper/lower)case encoding in the URI.
	http.HandleFunc("/link/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, content{
			ShowPaths: true,
			RawPath:   r.URL.RawPath,
			Path:      r.URL.Path,
		})
	})

	fmt.Println("Go to http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}