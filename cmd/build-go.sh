#!/bin/bash
# windows
GOOS=windows GOARCH=amd64 go build -o ./.bin/lucid-windows.exe ./cmd/serve/main.go

# darwin (mac os)
GOOS=darwin GOARCH=amd64 go build -o ./.bin/lucid-darwin ./cmd/serve/main.go

# linux
GOOS=linux GOARCH=amd64 go build -o ./.bin/lucid-linux ./cmd/serve/main.go

# default
go build -o ./.bin/lucid ./cmd/serve/main.go
