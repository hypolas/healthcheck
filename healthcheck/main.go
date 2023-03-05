package main

import (
	"flag"
	"os"

	helpers "github.com/hypolas/hypolashlckhelpers"
	http "github.com/hypolas/hypolashlckhttp"

)

func main() {
	// Run option
	customID := flag.String("id", "", "Needed for run chain of healthcheck")
	enableDebug := flag.Bool("debug", false, "Write debug variable in file (en var: HYPOLAS_LOGS_FILE)")
	flag.Parse()

	n, _ := os.Executable()
	log.Info.Println(n)

	// Config from flag to all modules
	os.Setenv("HYPOLAS_HEALTHCHECK_ID", *customID)
	if *enableDebug {
		os.Setenv("HYPOLAS_HEALTHCHECK_DEBUG", "true")
	}

	// Run healthcheck
	result := helpers.Result{}
	switch healthcheckType {
	case "http":
		result = http.Call()
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
