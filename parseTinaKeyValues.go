package main

import (
	"fmt"
	"os"
	"regexp"

	//	"slices"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// Define a new type for the enum
type Nesting int

// Define enum values using the const keyword and iota for sequential values
const (
    Root Nesting = iota
    Depth
    Close
)

type Queue struct {
	items []string
}

type CaptureZone struct {
	Start int
	Stop  int
}

type Zones struct {
	Zone []CaptureZone
}

var parseValueAsMarkdown bool

var printAst bool
var fileRegex = regexp.MustCompile(`(---)[^===][a-zA-Z1-9]*(---)`)
var keyRegex = regexp.MustCompile(`(\B- {1}|\b)(?:[a-z]([^'https]|\n)[a-zA-Z]+:)`)

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

func toString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func parseTinaKeyValues(file string) map[string]interface{} {
	keyQueue := new(Queue)
	valueQueue := new(Queue)
	keySlice := keyRegex.FindAllString(file, -1)
	keyIndexes := keyRegex.FindAllStringIndex(file, -1)
	zones := createCaptureZones(keyIndexes)

	for _, key := range keySlice {
		k := strings.TrimSuffix(key, ":")
		keyQueue.Enqueue(k)

	}

	for _, zone := range zones {
		value := captureText(zone.Start+1, zone.Stop-1, file)
		if value != "\n" {
			valueQueue.Enqueue(value)
		}

	}
	fields := make(map[string]interface{})

    //nested := make(map[string]interface{})

	for i := 0; i < keyQueue.Length(); {
		v := valueQueue.Dequeue()
		s := toString(&v)
		k := keyQueue.Dequeue()
        //parent := ""
		//n := Root
		fields[k] = s
	
	}
	fmt.Println(fields)
	return fields
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
