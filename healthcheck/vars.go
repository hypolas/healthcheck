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
	healthcheckHttpExpected string

	// JsonPath Flatter with double _
	healthcheckHttpJsonPath string

	// URL to check
	healthcheckHttpUrl string

	// Proxy if needed
	healthcheckHttpProxy string

	// Add header if needed
	healthcheckHttpHeaders string

	// Use return code ?
	healthcheckHttpUseCode bool

	// Define HTTP Timeout
	healthcheckHttpTimeout time.Duration

	// Check HTTP Status Code
	healthcheckHttpResponse []int

	returnedValue string
	separator     = "__"
	isJsonEntry   = true

	err error
)

type JsonKey struct {
	Name       string
	KeyIsArray bool
	ArrayIndex int
	IsLast     bool
}
