package main

import (
	"fmt"
	"os"
	"regexp"
	//"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Stack struct {
	items []interface{}
}

type KeyValue struct {
	Key   string
	Value string
}

type CaptureZone struct {
    Start   int
    Stop    int
}

type Zones struct {
    Zone []CaptureZone
}

var tokens []string
var parseStart bool
var parseEnd bool
var key string
var parseValueAsMarkdown bool
var value string
var keyValues []KeyValue
var printAst = false
var fileRegex = regexp.MustCompile(`(---)[^===][a-zA-Z1-9]*(---)`)
var keyRegex = regexp.MustCompile(`(\B- {1}|\b)(?:[a-z]([^'https]|\n)[a-zA-Z]+:)`)
var nestedRegex = regexp.MustCompile(`^[-]`)
var textValueRegex = regexp.MustCompile(`\>`)

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func parseTinaKeyValues(slicedBytes []string, file string) []KeyValue {
	keyStack := new(Stack)
    valueStack := new(Stack)
	//toParse := strings.Join(slicedBytes, " ")
	keySlice := keyRegex.FindAllString(file, -1)
    keyIndexes := keyRegex.FindAllStringIndex(file, -1)
    zones := createCaptureZones(keyIndexes)

	keyValue := KeyValue{
		Key:   key,
		Value: value,
	}
	keyValues = append(keyValues, keyValue)
	key = ""
	for _, key := range keySlice {
		// fmt.Println(key)
		keyStack.Push(key)

	}

    for _, zone := range zones {
        value := captureText(zone[0], zone[1], slicedBytes)
        valueStack.Push(value)
    }

	return keyValues
}

func createCaptureZones(keyIndexes [][]int) [][]int {
     
    textZones := make(Zones)
    
    for i:=0; i < len(keyIndexes);i = i + 2{
        textCaptureZone := CaptureZone{ 
            Start:keyIndexes[i][1],
            Stop:keyIndexes[i][0],
        }
        textZones = append(textZones, textCaptureZone)
    }

    return 
}


func captureText(start int, stop int, slicedBytes []string) string {
	text := ""
	for i := start; i < stop; i++ {
			text = text + slicedBytes[i] + " "
	}
	return text

}

func printStartEnd() {
	if !parseStart {
		parseStart = true
		fmt.Println("ParseStart", parseStart)
	} else {
		parseEnd = true
		fmt.Println("ParseEnd", parseEnd)

	}
}

func appendKeyValue(key string, value string) {
	if parseValueAsMarkdown {
		value = string(mdToHTML([]byte(value)))
	}

	keyValue := KeyValue{
		Key:   key,
		Value: value,
	}
	keyValues = append(keyValues, keyValue)
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	if printAst {
		fmt.Print("--- AST tree:\n")
		ast.Print(os.Stdout, doc)
		fmt.Print("\n")
	}

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
