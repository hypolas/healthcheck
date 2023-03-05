export HYPOLAS_HEALTHCHECK_TYPE=http
export HYPOLAS_HEALTHCHECK_HTTP_JSON=headers__Accept-Encoding
export HYPOLAS_HEALTHCHECK_HTTP_URL=https://httpbin.org/get
export HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=gzip
export HYPOLAS_LOGS_FILE=test/logs.log

go run . -debug

export HYPOLAS_HEALTHCHECK_MYCFG_TYPE=http
export HYPOLAS_HEALTHCHECK_MYCFG_HTTP_JSON=headers__Accept-Encoding
export HYPOLAS_HEALTHCHECK_MYCFG_HTTP_URL=https://httpbin.org/get
export HYPOLAS_HEALTHCHECK_MYCFG_HTTP_EXPECTED=gzip
export HYPOLAS_LOGS_FILE=test/logs.log

go run . -id MYCFG -debug



go test .