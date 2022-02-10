#!/bin/bash
go build ./craft.go && mv -f ./craft ./.build/craft
go build ./serve.go && mv -f ./serve ./.build/serve
