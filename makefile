#!/bin/bash
.check_binaries:
	./cmd/check-binaries.sh
.build_go:
	./cmd/build-go.sh
.build_svelte:
	./cmd/build-svelte.sh

build: .check_binaries .build_go .build_svelte
