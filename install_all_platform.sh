#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pass-gener_linux  pass-gener.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o pass-gener_windows.exe pass-gener.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o pass-gener_darwin pass-gener.go

CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o pass-gener_linux_arm pass-gener.go