package main

import (
	"fmt"
    "bufio"
    "os"
    "strings"

//	"github.com/gomarkdown/markdown"
//	"github.com/gomarkdown/markdown/html"
//	"github.com/gomarkdown/markdown/parser"
)

type pageFromMd struct {
    file string
    heroTitle string
    subtitle string
    aboutTitle string
    aboutBioMD string
    projectsTitle string
    projects []project
    blogTitle string
    blogPosts []blogPost
}

type project struct {
    title string
    descriptionMD string
    siteLink string
    githubLink string
    image string
}

type blogPost struct {
    title string
    date string
    contentMD string
}

func parseMdToHtml(filePath string) string {
    // Read the file contents
    file, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Error opening file: %s", err)
    }
    scanner := bufio.NewScanner(strings.NewReader(string(file)))
    split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		return
	}

    scanner.Split(split)
   	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	} 

	// Define the regex pattern for matching keys
    return scanner.Text()
}
