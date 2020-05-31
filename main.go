package main

import (
	// "log"
	// "net/http"
	"fmt"
	"log"
	"net/http"

	"github.com/evankim20/cyoa/parse"
)

func main() {
	story := parse.ParseJSON()
	for k := range story {
		fmt.Println(k)
	}

	fmt.Println("Starting server")

	// http.HandleFunc("/", makeHandler(viewHandler))
	log.Fatal(http.ListenAndServe(":8080", parse.HandlePage(story)))
}
