package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func getHTTP() {
	clientHTTP := constructHttpClient()
	prinfDebug(clientHTTP, "clientHTTP")

	reqHTTP, err := http.NewRequest("GET", healthcheckHTTPUrl, nil)
	prinfDebug(reqHTTP, "reqHTTP")
	printErr(err)
	reqHTTP.Header.Add("Accept", `application/json`)

	additionnalHeaders := splitFlatten(healthcheckHTTPHeaders)
	for _, header := range additionnalHeaders {
		splitedHeader := strings.Split(header, ",")
		reqHTTP.Header.Add(splitedHeader[0], splitedHeader[1])
	}

	resp, err := clientHTTP.Do(reqHTTP)

	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	defer resp.Body.Close()

	bodyHTTP, err := ioutil.ReadAll(resp.Body)
	prinfDebug(bodyHTTP, "bodyHTTP")
	printErr(err)

	/*
	*	If chek is on status html code the test stop here
	 */
	prinfDebug(healthcheckHTTPUseCode, "healthcheckHttpUseCode")
	if healthcheckHTTPUseCode {
		if intIsIn(resp.StatusCode, healthcheckHTTPResponse) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	/*
	*	If chek is on REST API, the json will be tested
	 */
	prinfDebug(healthcheckHTTPJsonPath, "healthcheckHttpJsonPath")
	if healthcheckHTTPJsonPath != "" {
		log.Println("taskJson")
		taskJSON(bodyHTTP)
	} else {
		returnedValue = strings.Trim(string(bodyHTTP), "\"")
	}
}

/*
*	Construct client HTTP
 */
func constructHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   0,
	}

	if healthcheckHTTPProxy != "" {
		proxyURL, err := url.Parse(healthcheckHTTPProxy)
		printErr(err)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	if healthcheckHTTPTimeout != 0 {
		client.Timeout = healthcheckHTTPTimeout * time.Second
	}

	return client
}

func intIsIn(i int, arrayInt []int) bool {
	for _, inte := range arrayInt {
		if i == inte {
			return true
		}
	}
	return false
}

func printErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
