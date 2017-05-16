/*
  After running this gif
      https://67.media.tumblr.com//b02659b27081314a412215f4eb31dacf/tumblr_nu2x4whJgy1udf6d3o1_400.gif
  through ffmpeg, Go fails to decode the gif with error:
      `gif: too much image data`

  Command:
  $ ffmpeg -i https://67.media.tumblr.com//b02659b27081314a412215f4eb31dacf/tumblr_nu2x4whJgy1udf6d3o1_400.gif outf.gif
  $ go run main.go --source outf.gif
  - Repro URL
  http://ualberta.ca/~odeke/tx.gif

  Interestingly:
  - The outfile is 1.0MB.
  - The original file is 1.4MB.
*/
package main

import (
    "flag"
    "image/gif"
    "log"
    "net/http"
    "path/filepath"
    "strings"
)

func localAndNetBasedClient() *http.Client {
    transport := new(http.Transport)
    transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
    client := new(http.Client)
    client.Transport = transport
    return client
}

var httpPrefixes = []string{"http", "https"}

func ensureFitForFetch(source string) string {
    for _, prefix := range httpPrefixes {
        if strings.HasPrefix(source, prefix) {
            return source
        }
    }

    absSource, err := filepath.Abs(source)
    if err == nil {
        source = absSource
    }
    return strings.Join([]string{"file://", source}, "/")
}

func main() {
    var source string
    flag.StringVar(&source, "source",
        "https://67.media.tumblr.com//b02659b27081314a412215f4eb31dacf/tumblr_nu2x4whJgy1udf6d3o1_400.gif",
        "the URI of the gif you want to ensure properly decodes in Go")
    flag.Parse()

    client := localAndNetBasedClient()
    source = ensureFitForFetch(source)
    res, err := client.Get(source)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()

    if _, err := gif.DecodeAll(res.Body); err != nil {
        log.Fatal(err)
    }
}