package main

import (
	"fmt"
	"net/http"

	"github.com/nao18/docker-golang-webapp01/src/pkg/handlers"
)

const posrNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", posrNumber))
	http.ListenAndServe(posrNumber, nil)
}
