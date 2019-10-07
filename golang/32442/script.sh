#!/bin/bash

GOOS=js GOARCH=wasm go test -c -exec=$GOPATH/src/go.googlesource.com/go/misc/wasm/go_js_wasm_exec -o main.wasm
GOOS=js GOARCH=wasm go tool test2json -t ./main.wasm -test.v -test.run '^Test_main$'
