#!/usr/bin/env bash

export HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=down
export HYPOLAS_HEALTHCHECK_HTTP_PROXY=
export HYPOLAS_HEALTHCHECK_HTTP_URL=https://httpbin.org/get
export HYPOLAS_HEALTHCHECK_TYPE=http
export HYPOLAS_HEALTHCHECK_HTTP_HEADERS=firstHeader,firstValue__Authorization,"Basic dGVzdDp0ZXN0MTIz"
export HYPOLAS_HEALTHCHECK_HTTP_JSON=headers__Accept-Encoding
