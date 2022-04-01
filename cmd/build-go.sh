#!/bin/bash
# windows
GOOS=windows GOARCH=amd64 go build -o ./.build/gorvel-windows.exe ./cmd/serve/main.go

# darwin (mac os)
GOOS=darwin GOARCH=amd64 go build -o ./.build/gorvel-darwin ./cmd/serve/main.go

# linux
GOOS=linux GOARCH=amd64 go build -o ./.build/gorvel-linux ./cmd/serve/main.go
