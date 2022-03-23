#!/bin/bash
build_gorvel:
	go build -o gopher.compiled ./cmd/gopher/main.go && \
		  mv -f gopher.compiled ./.build/gopher
	go build -o serve.compiled ./cmd/serve/main.go && \
		  mv -f serve.compiled ./.build/serve
