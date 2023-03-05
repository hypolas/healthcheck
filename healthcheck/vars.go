package main

import (
	"os"

	logg "github.com/hypolas/hypolaslogger"
)

var (
	log                     = logg.NewLogger("")
	healthcheckHTTPExpected string
	healthcheckType         = os.Getenv("HYPOLAS_HEALTHCHECK_TYPE")
)
