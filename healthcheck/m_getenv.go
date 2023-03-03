package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func taskLoadEnvironnement() {
	healthcheckType = os.Getenv("HYPOLAS_HEALTHCHECK_TYPE")

	/*
	*	Http check variable
	 */
	healthcheckHttpExpected = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_EXPECTED", ""))
	healthcheckHttpJsonPath = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_JSON", ""))
	healthcheckHttpUrl = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_URL", ""))
	healthcheckHttpProxy = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_PROXY", ""))
	healthcheckHttpHeaders = resolveVariable(getEnv("HYPOLAS_HEALTHCHECK_HTTP_HEADERS", ""))

	healthcheckHttpTimeout, err = time.ParseDuration(getEnv("HYPOLAS_HEALTHCHECK_HTTP_TIMEOUT", "0") + "s")
	printErr(err)

	statusCode := strings.Split(getEnv("HYPOLAS_HEALTHCHECK_HTTP_RESPONSES", ""), ",")
	if statusCode[0] != "" {
		healthcheckHttpUseCode = true
		for _, status := range statusCode {
			code, err := strconv.Atoi(status)
			printErr(err)
			healthcheckHttpResponse = append(healthcheckHttpResponse, code)
		}
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
