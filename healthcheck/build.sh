#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-w -h -H windowsgui -extldflags=-static" -o bin/healthcheck_hpl-win-amd64.exe .
GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -h -extldflags=-static" -o bin/healthcheck_hpl-linux-amd64 .
(
    cd bin || return
    cp healthcheck_hpl-win-amd64.exe healthcheck_hpl-win-amd64-upx-compress.exe
    cp healthcheck_hpl-linux-amd64 healthcheck_hpl-linux-amd64-upx-compress
    upx --ultra-brute --lzma ./*upx-compress*
)
