#!/bin/bash
go build -o serve.compiled ./cmd/serve/main.go && \
        mv -f serve.compiled ./.build/gorvel
