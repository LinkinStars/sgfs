#!/usr/bin/env bash

# mac os
# go build ./main/sgfs.go

# linux 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -s -a -installsuffix cgo ./main/sgfs.go

# tar
# tar -zcvf sgfs-1.0.2-linux.tar.gz sgfs-1.0.2
