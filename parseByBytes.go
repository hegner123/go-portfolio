package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Tina struct {
	Key   string
	Value string
}

var toParsePair string = "---"
var tokenStart string
var tokenEnd string
var token string
var hasKey bool
var hasValue bool
var currentKey string
var currentVal string
var tokenStatus string = "stop"
var tokenList []string

func parseMdByBytes(filePath string) (string){
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

    stringFile := string(file)

	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		parseBytes(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return stringFile
}

func parseBytes(tokenArg string) {
	// validate tokenArg is a valid unicode
	isValidUnicode := utf8.ValidString(tokenArg)
	if isValidUnicode {
		handleValidUnicode(tokenArg)
	}
}

func handleValidUnicode(char string) {
	switch char {
	// check if char is a whitespace character
	case " ", "\t", "\n":
        if len(token) > 0 {
			handleTokenSeperator()
        }
	default:
        token = token + char
	}
}

func handleTokenSeperator() {
	tokenList = append(tokenList, token)	
	token = ""
}
