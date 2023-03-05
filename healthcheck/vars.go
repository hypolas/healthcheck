package main

import (
	"os"

	helpers "github.com/hypolas/hypolashlckhelpers"

)

var (
	log                     = helpers.NewLogger()
	healthcheckHTTPExpected string
	healthcheckType         = os.Getenv("HYPOLAS_HEALTHCHECK_TYPE")
)
