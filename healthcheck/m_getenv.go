package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

/*
*	Get entry environnement variables and parse it
 */
func taskLoadEnvironnement() {
	healthcheckType = os.Getenv("HYPOLAS_HEALTHCHECK_TYPE")
	healthcheckLogsFolder = getEnv("HYPOLAS_HEALTHCHECK_LOGS_FOLDER", "logs/")
	os.Setenv("HYPOLAS_LOGS_FILE", healthcheckLogsFolder)

	/*
	*	Http check variable
	 */
	healthcheckHTTPExpected = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_EXPECTED", ""))
	healthcheckHTTPJsonPath = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_JSON", ""))
	healthcheckHTTPUrl = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_URL", ""))
	healthcheckHTTPProxy = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_PROXY", ""))
	healthcheckHTTPHeaders = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_HEADERS", ""))

	healthcheckHTTPTimeout, err = time.ParseDuration(getEnv("HYPOLAS_HEALTHCHECK_HTTP_TIMEOUT", "0") + "s")
	printErr(err)

	statusCode := strings.Split(getEnv("HYPOLAS_HEALTHCHECK_HTTP_RESPONSES", ""), ",")
	if statusCode[0] != "" {
		healthcheckHTTPUseCode = true
		for _, status := range statusCode {
			code, err := strconv.Atoi(status)
			printErr(err)
			healthcheckHTTPResponse = append(healthcheckHTTPResponse, code)
		}
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
