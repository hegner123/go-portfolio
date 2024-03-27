package main

import (
    "fmt"
    "log"
    "os"
//    "net/http"
//    "strconv"
//    "strings"
//    "github.com/a-h/templ"
)



func defineRoutes() {
	files, err := os.ReadDir("./content/pages")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}


}
