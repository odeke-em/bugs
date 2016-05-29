package main

import (
    pb "github.com/dvyukov/go-fuzz/examples/protobuf/pb"
    "github.com/golang/protobuf/proto"
)

func main() {
    data := []byte("\n\x02\n\x00")
    v := new(pb.M24)
    err := proto.Unmarshal(data, v)
    if err != nil {
        return
    }
    _, err = proto.Marshal(v)
    if err != nil {
        panic(err)
    }
}