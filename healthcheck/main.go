package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func init() {
	taskLoadEnvironnement()
}

func main() {
	debug = flag.Bool("d", false, "Debug with log print")
	flag.Parse()
	debugEnable = *debug

	switch healthcheckType {
	case "http":
		getHttp()
	}

	prinfDebug(returnedValue, "returnedValue")
	prinfDebug(healthcheckHttpExpected, "healthcheckHttpExpected")
	if returnedValue == healthcheckHttpExpected {
		log.Println("OK")
	} else {
		os.Exit(1)
	}
}

func splitFlatten(flatten string) []string {
	flatten = strings.TrimSpace(flatten)
	if !strings.Contains(flatten, separator) {
		return []string{}
	}
	return strings.Split(flatten, separator)
}

/*
* 	Print debug if enabled
 */
func prinfDebug(info interface{}, name string) {
	if debugEnable {
		log.Println(strings.Repeat("#", 20))
		log.Printf("Name: %s\n", name)
		varType := fmt.Sprintf("%+v\n", reflect.TypeOf(info))
		log.Printf("Type: %s", varType)

		if varType == "[]uint8" {
			varValue := string(reflect.ValueOf(info).Bytes())
			log.Printf("Value: %s", varValue)
		} else {
			log.Printf("Value: %s", reflect.ValueOf(info))
		}
		log.Println(strings.Repeat("#", 20))
	}
}
