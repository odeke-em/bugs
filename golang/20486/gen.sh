#!/bin/bash
go tool compile lib.go
go tool compile main.go
go tool link main.o
