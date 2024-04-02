package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"os"
	"regexp"
)

type Stack struct {
    items []interface{}
}

type KeyValue struct {
	Key   string
	Value string
}

var tokens []string
var parseStart bool
var parseEnd bool
var key string
var parseValueAsMarkdown bool
var value string
var keyValues []KeyValue
var printAst = false
var keyRegex = regexp.MustCompile(`([^\'|https][a-z])([a-zA-Z][^:])+:`)
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

func parseTinaKeyValues(slicedBytes []string) []KeyValue {
    tokenStack := new( Stack)
	for i := 0; i < len(slicedBytes); i++ {
        token := slicedBytes[i]
		if token == "---" {
			//handle nested operations

		} else if keyRegex.MatchString(token)  {
            tokenStack.Push(token)
		} else if textValueRegex.MatchString(token) {
			fmt.Println("text value capture")
			textValue, l := captureText(i+1, slicedBytes, tokenStack)
			i = l
			fmt.Println(textValue)
			fmt.Println(l)
		} else if nestedRegex.MatchString(token) {
			fmt.Println("nestedRegex")
			fmt.Println(token)
			//fmt.Println(token)
		}
	}
	return keyValues
}

func haveFun(i int, slicedBytes []string, s *Stack){
    value, l := captureText(i+1, slicedBytes, s)
    fmt.Println(l)
    keyValue := KeyValue{
        Key:   key,
        Value: value,
    }
    fmt.Println(keyValue)
    keyValues = append(keyValues, keyValue)
    key = ""

}

func captureText(iter int, slicedBytes []string, s *Stack) (string, int) {
	text := ""
	it := 0
	for i := iter; i < len(slicedBytes); i++ {
		if !keyRegex.MatchString(slicedBytes[i]) {
			text = text + slicedBytes[i] + " "
		} else {
			it = i
		}
	}
	return text, it

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
