package main

import (
    "encoding/gob"
    "fmt"
    "io/ioutil"
    "log"
)

type File struct {
    A, B string
    V    int
}

func main() {
    f, err := ioutil.TempFile("", "foo")
    if err != nil {
        panic(err)
    }

    e := gob.NewEncoder(f)

    m := make(map[string][]*File)
    for i := 0; i < (1 << 22); i += 1 {
        fl := File{"00000000000000000000000000000000000000000000000",
            "11111111111111111111111111111111111111111", 42}
        var arr []*File
        for j := 0; j < 4; j += 1 {
            arr = append(arr, &fl)
        }
        m[fmt.Sprintf("%d", i)] = arr
        if i%(1<<20) == 0 {
            log.Printf("i = %d / %d", i, 1<<22)
        }
    }

    if err := e.Encode(m); err != nil {
        panic(err)
    }
}