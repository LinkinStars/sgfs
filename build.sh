#!/usr/bin/env bash

# mac os
# go build ./main/sgfs.go

# linux 
# GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -s -a -installsuffix cgo ./main/sgfs.go

# tar
# tar -zcvf sgfs-1.0.2-linux.tar.gz sgfs-1.0.2

version=$(git describe --tags)

echo "build darwin"
go build -o ./bin/darwin/sgfs ./main/sgfs.go

echo "build linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -installsuffix cgo  -o ./bin/linux/sgfs ./main/sgfs.go

echo "build windows"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/windows/sgfs.exe ./main/sgfs.go

echo ./bin/darwin ./bin/linux ./bin/windows | xargs -n 1 cp -v ./conf.yml
echo ./bin/darwin ./bin/linux | xargs -n 1 cp -v ./shutdown.sh
echo ./bin/darwin ./bin/linux | xargs -n 1 cp -v ./startup.sh

tar -zcvf ./bin/sgfs-$version-darwin.tar.gz -C ./bin/darwin .
tar -zcvf ./bin/sgfs-$version-linux.tar.gz -C ./bin/linux .
tar -zcvf ./bin/sgfs-$version-windows.tar -C ./bin/windows .
