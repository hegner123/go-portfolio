package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
    fmt.Println("--------------------------------------------------------------------------------------------------")
    content := "./content/pages/home.md";

    p := parseTinaMd(content)
   //fmt.Println(tinaList) 
    
	component := pageFromMarkdown(p)
	
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
    http.ListenAndServe(":3000", nil)

}



