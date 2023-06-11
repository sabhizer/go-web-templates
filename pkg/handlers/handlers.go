package handlers

import (
	"net/http"
	"web-templates/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Home Page")
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
