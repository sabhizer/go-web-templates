package main

import (
	"fmt"
	"net/http"
	"web-templates/pkg/handlers"
)

var port = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting server on port: ", port)
	http.ListenAndServe(port, nil)
}
