package main

import (
	"os"

	logg "github.com/hypolas/hypolaslogger"
)

var (
	logf                    = logg.NewLogger("")
	returnedValue           string
	healthcheckHTTPExpected string
	healthcheckType         = os.Getenv("HYPOLAS_HEALTHCHECK_TYPE")
)
