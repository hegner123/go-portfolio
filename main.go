package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	printGreen("------------------------------------------------------------")
	page := fetchPage()

	component := portfolio(page)
	http.Handle("/", templ.Handler(component))
	//http.Handle("/blog", templ.Handler(blogArchive()))

	printCyan("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		printRedErr(err)
	}

}
