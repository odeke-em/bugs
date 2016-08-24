package main

import (
  "bytes"
)

func main() {
  _ = bytes.Repeat([]byte("a"), 1.8e8) // int(^uint(0)>>26))
}
