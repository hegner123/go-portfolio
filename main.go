package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

    content := "./content/pages/home.md";
    html := parseMdToHtml(content) 
      //bytes to string
    stringHtml := string(html)
    

	component := pageFromMarkdown(stringHtml)
	
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}



