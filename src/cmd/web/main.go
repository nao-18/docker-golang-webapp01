package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nao18/docker-golang-webapp01/src/pkg/config"
	"github.com/nao18/docker-golang-webapp01/src/pkg/handlers"
	"github.com/nao18/docker-golang-webapp01/src/pkg/render"
)

const posrNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", posrNumber))
	_ = http.ListenAndServe(posrNumber, nil)
}
