#!/usr/bin/env bash
# build: files-extract-tool

env GOOS=linux GOARCH=amd64 go build -o ./bin/files-extract-tool main.go
