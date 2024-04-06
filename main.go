package main

import (
	"fmt"
	"net/http"
    "github.com/a-h/templ"
)



func main() {
	fmt.Println("--------------------------------------------------------------------------------------------------")
    page := fetchPage()


	component := portfolio(page)
	http.Handle("/", templ.Handler(component))
    //http.Handle("/blog", templ.Handler(blogArchive()))

	fmt.Println("Listening on :3000")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        fmt.Println(err)
    }

}
