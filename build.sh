#/usr/bin/env bash
# build: files-extraction-tool

env GOOS=linux GOARCH=amd64 go build -o ./bin/files-extraction-tool main.go