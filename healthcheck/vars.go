package main

import (
	"time"
)

var (
	// Type de check (http, tcp)
	healcheckType string

	// Expected http value
	healcheckHttpExpected string

	// JsonPath Flatter with double _
	healcheckHttpJsonPath string

	// URL to check
	healcheckHttpUrl string

	// Proxy if needed
	healcheckHttpProxy string

	// Add header if needed
	healcheckHttpHeaders string

	// Define HTTP Timeout
	healcheckHttpTimeout time.Duration

	// Check HTTP Status Code
	healcheckHttpResponse []int

	returnedValue string
	separator     = "__"
	isJsonEntry   = true
)

type JsonKey struct {
	Name       string
	KeyIsArray bool
	ArrayIndex int
	IsLast     bool
}
