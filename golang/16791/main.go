package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func main() {
    f, _ := os.Open("mock_data.csv")
    defer f.Close()

    r := csv.NewReader(bufio.NewReader(f))
    for {
        line, err := r.Read()
        if err == io.EOF {
            break
        }
        if line[0] == "42" {
            fmt.Println(line)
        }
    }

}
