package main

import (
	"encoding/json"
	"strconv"
	"strings"
)

/*
*
*	Parse flatten JSON and return expected value
*
 */

func taskJSON(testJSON []byte) {
	jsonPathDecomposer(healthcheckHTTPJsonPath, testJSON)
}

/*
*
*	Read flatten key
*
 */
func jsonPathDecomposer(jpath string, jsonFile []byte) {
	arrayPath := splitFlatten(jpath)
	lenPath := len(arrayPath)
	var skipThis = 99999
	for i, jp := range arrayPath {
		if skipThis == i {
			continue
		}
		jsonDef := keyTypeDecomposer(jp, i, lenPath > i+1, arrayPath)

		if lenPath == i+1 {
			jsonDef.IsLast = true
		}

		if jsonDef.KeyIsArray {
			skipThis = i + 1
		}
		jsonFile = jsonDecomposer(jsonDef, jsonFile)
	}
	returnedValue = strings.Trim(string(jsonFile), "\"")
}

func jsonDecomposer(jsonFormat JSONKey, jsonFile []byte) []byte {
	var inner interface{}

	if jsonFormat.Name == "" {
		return jsonFile
	}

	switch jsonFormat.KeyIsArray {
	case true:
		theInterface := map[string][]interface{}{}
		json.Unmarshal(jsonFile, &theInterface)
		inner = theInterface[jsonFormat.Name][jsonFormat.ArrayIndex]
	case false:
		theInterface := map[string]interface{}{}
		json.Unmarshal(jsonFile, &theInterface)
		inner = theInterface[jsonFormat.Name]
	}

	jsonInner, err := json.Marshal(inner)
	prinfDebug(jsonInner, "jsonInner")
	printErr(err)

	return jsonInner
}

func keyTypeDecomposer(key string, index int, haveNext bool, arrayPath []string) JSONKey {
	tmpKey := JSONKey{}
	if haveNext {
		if ind, err := strconv.Atoi(arrayPath[index+1]); err == nil {
			tmpKey.KeyIsArray = true
			tmpKey.ArrayIndex = ind
		}
	}

	tmpKey.Name = key

	prinfDebug(tmpKey, "tmpKey")

	return tmpKey
}

func isIn(s string, array []string) bool {
	for _, a := range array {
		if a == s {
			return true
		}
	}
	return false
}
