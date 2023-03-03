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
	client := constructHttpClient()

	req, _ := http.NewRequest("GET", healcheckHttpUrl, nil)
	req.Header.Add("Accept", `application/json`)

	additionnalHeaders := splitFlatten(healcheckHttpHeaders)
	for _, header := range additionnalHeaders {
		splitedHeader := strings.Split(header, ",")
		req.Header.Add(splitedHeader[0], splitedHeader[1])
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if intIsIn(resp.StatusCode, healcheckHttpResponse) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
	log.Println(resp.StatusCode)

	if healcheckHttpJsonPath != "" {
		log.Println("taskJson")
		taskJson(body)
	}
}

func constructHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   0,
	}

	if healcheckHttpProxy != "" {
		proxyUrl, _ := url.Parse(healcheckHttpProxy)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}

	if healcheckHttpTimeout != 0 {
		client.Timeout = healcheckHttpTimeout * time.Second
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
