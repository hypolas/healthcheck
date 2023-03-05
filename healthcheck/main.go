package main

import (
	"os"

	http "github.com/hypolas/hypolashlckhttp"
)

func main() {
	switch healthcheckType {
	case "http":
		returnedValue = http.GetHTTP()
	}

	logf.VarDebug(returnedValue, "returnedValue")
	logf.VarDebug(healthcheckHTTPExpected, "healthcheckHttpExpected")
	if returnedValue == healthcheckHTTPExpected {
		logf.Info.Println("OK")
	} else {
		os.Exit(1)
	}
}
