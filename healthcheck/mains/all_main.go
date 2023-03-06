package main

import (
	"flag"
	"os"

	helpers "github.com/hypolas/hypolashlckhelpers"
    mhlck_http "github.com/hypolas/hypolashlckhttp"
    mhlck_ping "github.com/hypolas/hypolashlckping"
)

func main() {
	// Run option
	customID := flag.String("id", "", "Needed for run chain of healthcheck")
	enableDebug := flag.Bool("debug", false, "Write debug variable in file (en var: HYPOLAS_LOGS_FILE)")
	flag.Parse()

	// Config from flag to all modules
	os.Setenv("HYPOLAS_HEALTHCHECK_ID", *customID)
	if *enableDebug {
		os.Setenv("HYPOLAS_HEALTHCHECK_DEBUG", "true")
	}

	// Run healthcheck
	result := helpers.Result{}
	switch healthcheckType {
	case "http":
		result = mhlck_http.Call()
	case "ping":
		result = mhlck_ping.Call()
	}


	log.VarDebug(result, "result")
	log.VarDebug(healthcheckHTTPExpected, "healthcheckHttpExpected")
	if result.IsUP {
		log.Info.Println("OK")
	} else {
		log.Err.Println("KO")
		os.Exit(1)
	}
}
