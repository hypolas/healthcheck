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

func getHttp() {
	clientHTTP := constructHttpClient()
	prinfDebug(clientHTTP, "clientHTTP")

	reqHTTP, err := http.NewRequest("GET", healthcheckHttpUrl, nil)
	prinfDebug(reqHTTP, "reqHTTP")
	printErr(err)
	reqHTTP.Header.Add("Accept", `application/json`)

	additionnalHeaders := splitFlatten(healthcheckHttpHeaders)
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

	prinfDebug(healthcheckHttpUseCode, "healthcheckHttpUseCode")
	if healthcheckHttpUseCode {
		if intIsIn(resp.StatusCode, healthcheckHttpResponse) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	prinfDebug(healthcheckHttpJsonPath, "healthcheckHttpJsonPath")
	if healthcheckHttpJsonPath != "" {
		log.Println("taskJson")
		taskJson(bodyHTTP)
	} else {
		returnedValue = strings.Trim(string(bodyHTTP), "\"")
	}
}

func constructHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   0,
	}

	if healthcheckHttpProxy != "" {
		proxyUrl, err := url.Parse(healthcheckHttpProxy)
		printErr(err)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}

	if healthcheckHttpTimeout != 0 {
		client.Timeout = healthcheckHttpTimeout * time.Second
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
