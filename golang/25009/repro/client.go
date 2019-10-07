package main

import (
        "bytes"
        "crypto/tls"
        "errors"
        "fmt"
        "io/ioutil"
        "net/http"
        "sync"
        "time"
)

type client struct {
        base string
        hc   *http.Client
}

func newClient(baseURL string) *client {
        return &client{
                base: baseURL,
                hc: &http.Client{
                        Timeout: 5 * time.Second,
                        Transport: &http.Transport{
                                MaxIdleConnsPerHost:   100,
                                MaxIdleConns:          100,
                                IdleConnTimeout:       90 * time.Second,
                                ExpectContinueTimeout: 5 * time.Second,
                                TLSClientConfig: &tls.Config{
                                        InsecureSkipVerify: true,
                                },
                        },
                },
        }
}

func (c *client) send(content []byte) error {
        if content == nil {
                return errors.New("nil content")
        }

        body := new(bytes.Buffer)
        body.Write(content)

        start := time.Now()
        req, err := http.NewRequest("POST", c.base, body)
        if err != nil {
                return fmt.Errorf("Send %+v (from %s)", err, time.Since(start))
        }
        req.Header = make(http.Header)
        req.Header.Set("Content-Type", "application/json")
        resp, err := c.hc.Do(req)
        if err != nil {
                return fmt.Errorf("Send: HTTPClient.Do: %+v (from %s)", err, time.Since(start))
        }
        defer resp.Body.Close()

        blob, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return fmt.Errorf("Send: Read: %+v (from %s)", err, time.Since(start))
        }
        fmt.Printf("Send: Resp: %s\n", blob)
        return nil
}

func main() {
        client := newClient("https://localhost:9789/")
        content := []byte(`{"test":1}`)
        var wg sync.WaitGroup
        for i := 0; i < 100; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        if err := client.send(content); err != nil {
                                fmt.Printf("#%d: err: %v\n", i, err)
                        }
                }(i)
        }
        wg.Wait()
}
