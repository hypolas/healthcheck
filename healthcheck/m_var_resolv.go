package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

/*
*	Format environnement variables
 */
func resolveVariable(strVar string) string {
	if strings.Contains(strVar, "#CMDSTART#") {
		strVar = resolveCMD(strVar)
	}
	outStr := os.ExpandEnv(strVar)
	prinfDebug(outStr, "outStr")

	return outStr
}

/*
*	Run cmd embeded in environnement variable and run it
 */
func resolveCMD(cmdString string) string {
	inputString := strings.TrimSpace(cmdString)
	strCommand := getStringInBetween(inputString, "#CMDSTART#", "#CMDEND#")
	stringToReplace := "#CMDSTART#" + strCommand + "#CMDEND#"
	prinfDebug(strCommand, "strCommand")
	prinfDebug(stringToReplace, "stringToReplace")

	cmdArgs := strings.Split(strings.TrimSpace(strCommand), " ")
	prinfDebug(cmdArgs, "cmdArgs")

	var cmd *exec.Cmd
	if len(cmdArgs) == 1 {
		cmd = exec.Command(cmdArgs[0])
	} else {
		cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
	}
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Env = os.Environ()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	/*
	*	Remove all unwanted characters
	*/
	cmdResult := strings.TrimFunc(stdout.String(), func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
	prinfDebug(cmdResult, "cmdResult")

	detectedCMDValue := strings.Replace(inputString, stringToReplace, cmdResult, -1)
	prinfDebug(detectedCMDValue, "detectedCMDValue")

	return detectedCMDValue
}

/*
*	Extract command from environnement variables
*/
func getStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return result
	}

	newS := str[s+len(start):]
	e := strings.Index(newS, end)
	if e == -1 {
		return result
	}
	result = newS[:e]
	return result
}
