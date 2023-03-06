#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-w -h -H windowsgui  -extldflags=-static" -o bin/healthcheck_hpl-win-amd64-$1.exe .
GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -h -extldflags=-static" -o bin/healthcheck_hpl-linux-amd64-$1 .
cp bin/healthcheck_hpl-win-amd64-$1.exe bin/healthcheck_hpl-win-amd64-$1-compressed.exe
cp bin/healthcheck_hpl-linux-amd64-$1 bin/healthcheck_hpl-linux-amd64-$1-compressed
