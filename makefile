#!/bin/bash
build:
	@echo "Check if go/npm/node exists"
	./cmd/check-binaries.sh

	@echo "Install github.com/cespare/reflex binary"
	./cmd/install-reflex.sh
