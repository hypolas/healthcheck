package main

import (
	"time"
)

var (
	// Debug
	debug       *bool
	debugEnable bool

	// Type de check (http, tcp)
	healthcheckType string

	// Expected http value
	healthcheckHTTPExpected string

	// JsonPath Flatter with double _
	healthcheckHTTPJsonPath string

	// URL to check
	healthcheckHTTPUrl string

	// Proxy if needed
	healthcheckHTTPProxy string

	// Add header if needed
	healthcheckHTTPHeaders string

	// Use return code ?
	healthcheckHTTPUseCode bool

	// Define HTTP Timeout
	healthcheckHTTPTimeout time.Duration

	// Check HTTP Status Code
	healthcheckHTTPResponse []int

	returnedValue string
	separator     = "__"
	isJSONEntry   = true

	err error
)

type JSONKey struct {
	Name       string
	KeyIsArray bool
	ArrayIndex int
	IsLast     bool
}
