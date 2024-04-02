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
	Start int
	Stop  int
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


type Queue struct {
    items []string
}

func (q *Queue) Enqueue(item string) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() string {
    if len(q.items) == 0 {
        return ""
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item
}

func (q *Queue) Length() int {
    return len(q.items)
}

func parseTinaKeyValues(slicedBytes []string, file string) []KeyValue {
	keyQueue := new(Queue)
	valueQueue := new(Queue)
	//toParse := strings.Join(slicedBytes, " ")
	keySlice := keyRegex.FindAllString(file, -1)
	keyIndexes := keyRegex.FindAllStringIndex(file, -1)
	zones := createCaptureZones(keyIndexes)

	for _, key := range keySlice {
		// fmt.Println(key)
		keyQueue.Enqueue(key)

	}

	for _, zone := range zones {
        value := captureText(zone.Start+1, zone.Stop-1, file)
        if value != "\n"{
        valueQueue.Enqueue(value)
        }

	}

    for i:=0; i<keyQueue.Length();{
        keyValue := KeyValue{
            Key: keyQueue.Dequeue(),
            Value: valueQueue.Dequeue(),
        }
        keyValues = append(keyValues,keyValue )
    }

	return keyValues
}

func createCaptureZones(keyIndexes [][]int) []CaptureZone {
	textZones := make([]CaptureZone, 0)
	for i := 1; i < len(keyIndexes); i++ {
		textCaptureZone := CaptureZone{
			Start: keyIndexes[i-1][1],
			Stop:  keyIndexes[i][0],
		}
		textZones = append(textZones, textCaptureZone)
	}
	return textZones
}

func captureText(start int, stop int, file string) string {
	text := ""
	for i := start; i < stop; i++ {
		text = text + string(file[i])
	}
	return text

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


