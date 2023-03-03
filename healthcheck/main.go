package main

import (
	"log"
	"os"
	"strings"
)

func init() {
	taskLoadEnvironnement()
}

func main() {
	switch healcheckType {
	case "http":
		getHttp()
	}

	log.Println(returnedValue)
	if returnedValue == healcheckHttpExpected {
		log.Println("OK")
	} else {
		os.Exit(1)
	}
}

func splitFlatten(flatten string) []string {
	return strings.Split(flatten, separator)
}
