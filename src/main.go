package main

import (
	"fmt"
	"net/http"
)

const posrNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", posrNumber))
	http.ListenAndServe(posrNumber, nil)
}
