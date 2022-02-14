#!/bin/bash
go build cmd/serve/main.go
mv main ./.build/main

# create the storage logs under .build/ folder
mkdir -p ./.build/storage/logs/
touch -f ./.build/storage/logs/gorvel.log
