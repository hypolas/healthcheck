#!/usr/bin/env bash
(
    cd bin || return
    upx --ultra-brute --lzma ./*-compressed-*
)
