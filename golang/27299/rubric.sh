#!/bin/sh

rm -rf newpack && mkdir newpack
cd newpack
GO111MODULE=on go mod init main
GO111MODULE=on go get github.com/antlr/antlr4/runtime/Go/antlr@v0.0.0-20180728001836-7d0787e29ca8
