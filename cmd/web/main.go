package main

import (
	"fmt"
	"net/http"
	"web-templates/pkg/config"
	"web-templates/pkg/handlers"
	"web-templates/pkg/render"
)

var port = ":8080"

func main() {
	tc, _ := render.CreateTemplateCache()
	fmt.Println("Template Cache:", tc)

	app := config.AppConfig{TemplateCache: tc}
	app.UseCache = true

	// Passing app config data/object from main package to render package for render package to be able to use it.
	render.NewTemplate(&app)

	// Passing app config data/object from main package to handler package as a new Repository struct.
	// Created local Repository struct in handlers package to convert Handler functions to Repository Reciever methods.
	handlers.NewRepository(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting server on port: ", port)
	http.ListenAndServe(port, nil)
}
