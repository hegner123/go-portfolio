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

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)

}
