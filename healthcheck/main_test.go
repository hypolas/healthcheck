package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// TestHTTPApi get API et test result
func TestMain(t *testing.T) {
	main()

	readFile, err := os.Open(log.LogFile.Name())

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
