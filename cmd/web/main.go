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

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting server on port: ", port)
	http.ListenAndServe(port, nil)
}
