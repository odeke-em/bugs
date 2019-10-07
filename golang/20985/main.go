package main

import (
    "flag"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "sync"

    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var (
    bucketName = flag.String("bucket", "", "bucket to use")
    objectName = flag.String("object", "", "object to read")
)

const n = 10

func main() {
    flag.Parse()
    ctx := context.Background()
    ts, err := google.DefaultTokenSource(ctx, "https://www.googleapis.com/auth/devstorage.read_only")
    if err != nil {
        log.Fatal(err)
    }
    hc := &http.Client{
        Transport: &oauth2.Transport{
            Source: ts,
            Base:   http.DefaultTransport,
        },
    }
    var wg sync.WaitGroup
    for i := 0; i < n; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            u := "https://storage.googleapis.com/" + *bucketName + "/" + *objectName
            req, err := http.NewRequest("GET", u, nil)
            if err != nil {
                log.Fatal(err)
            }
            req = req.WithContext(ctx)
            res, err := hc.Do(req)
            if err != nil {
                uerr, ok := err.(*url.Error)
                if ok {
                    log.Printf("url Error %#v, temporary=%v, timeout=%v",
                        uerr, uerr.Temporary(), uerr.Timeout())
                }
                log.Fatalf("NewReader: %v (%T)", err, err)
            }
            if res.StatusCode < 200 || res.StatusCode > 299 {
                body, _ := ioutil.ReadAll(res.Body)
                res.Body.Close()
                log.Fatal(res.StatusCode, res.Header, string(body))
            }
            defer res.Body.Close()
            if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
                log.Fatal(err)
            }
        }()
    }
    wg.Wait()
}
