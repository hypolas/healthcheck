package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func taskLoadEnvironnement() {
	healcheckType = os.Getenv("HYPOLAS_HEALCHECK_TYPE")

	/*
	*	Http check variable
	 */
	healcheckHttpExpected = os.Getenv("HYPOLAS_HEALCHECK_HTTP_EXPECTED")
	healcheckHttpJsonPath = os.Getenv("HYPOLAS_HEALCHECK_HTTP_JSON")
	healcheckHttpUrl = os.Getenv("HYPOLAS_HEALCHECK_HTTP_URL")
	healcheckHttpProxy = os.Getenv("HYPOLAS_HEALCHECK_HTTP_PROXY")
	healcheckHttpHeaders = os.Getenv("HYPOLAS_HEALCHECK_HTTP_HEADERS")

	healcheckHttpTimeout, _ = time.ParseDuration(os.Getenv("HYPOLAS_HEALCHECK_HTTP_TIMEOUT") + "s")

	statusCode := strings.Split(os.Getenv("HYPOLAS_HEALCHECK_HTTP_RESPONSES"), ",")
	for _, status := range statusCode {
		code, _ := strconv.Atoi(status)
		healcheckHttpResponse = append(healcheckHttpResponse, code)
	}
}
