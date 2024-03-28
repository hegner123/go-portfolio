package main

import (
    "os"
	"fmt"
    "regexp"
    "github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

)

type KeyValue struct {
	Key   string
	Value string
}

var key string
var parseValueAsMarkdown bool
var value string
var keyValues []KeyValue
var printAst = false


func parseTinaKeyValues(tokens []string) []KeyValue {
    // parseHeader := regexp.MustCompile("---")
	 keyRegex := regexp.MustCompile(`([^\'][a-z])([a-zA-Z][^:])+:`)
	// valueRegex := regexp.MustCompile(`: \b.+\n`)
    for _, token := range tokens {
        if token == "---" {
            continue
        }
        if token == ">" {
            parseValueAsMarkdown = true
            continue
        }
        if keyRegex.MatchString(token)  && key == "" {
            key = token
            
        } else if key != "" && keyRegex.MatchString(token) {
            appendKeyValue(key, value)
            key = token
            value = ""
            parseValueAsMarkdown = false
        } else {
            value += " " + token
        }
    }

    fmt.Println(keyValues[0].Key)	
    return keyValues
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
