#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pass-gener-linux  pass-gener.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o pass-gener-windows.exe pass-gener.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o pass-gener-darwin pass-gener.go

CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o pass-gener-linux-arm pass-gener.go