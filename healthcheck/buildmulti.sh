#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-w -h -H windowsgui  -extldflags=-static" -o bin/healthcheck_hpl-$1-win-amd64.exe .
GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -h -extldflags=-static" -o bin/healthcheck_hpl-$1-linux-amd64 .
cp bin/healthcheck_hpl-$1-win-amd64.exe bin/healthcheck_hpl-$1-compressed-win-amd64.exe
cp bin/healthcheck_hpl-$1-linux-amd64 bin/healthcheck_hpl-$1-compressed-linux-amd64
