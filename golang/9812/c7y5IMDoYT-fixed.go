package main

import (
    "encoding/json"
    "fmt"
    "time"
)

type A struct {
    Time time.Time `json:",string)"`
}

func main() {
    a := A{Time: time.Now()}
    b, _ := json.Marshal(&a)

    fmt.Println("Encoded:", string(b))

    err := json.Unmarshal(b, &a)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Decoded:%#v", a)

    // Golang expects quoted input - this succeeds:
    /*
        err := json.Unmarshal([]byte(`{"Time":"\"2009-11-10T23:00:00Z\""}`), &a)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Decoded:%#v", a)
    */
}
