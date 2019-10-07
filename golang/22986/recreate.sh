#!/bin/sh
mkdir -p 22986 && cd 22986 || exit "Failed to create 22986"
mkdir -p test/p-1.0 && echo "package foo\ntype Sample int" > test/p-1.0/foo.go
echo 'package test\n\nimport (\n"testing"\ntest "./test/p-1.0"\n)\nfunc TestFoo(t *testing.T) {\n_ = test.Sample(10)\n}' > m_test.go && go test -v
cd ../ && rm -rf 22986
