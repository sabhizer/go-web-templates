package handlers

import (
	"net/http"
	"web-templates/pkg/config"
	"web-templates/pkg/models"
	"web-templates/pkg/render"
)

type Respository struct {
	App *config.AppConfig
}

var Repo *Respository

func NewRepository(a *config.AppConfig) {
	Repo = &Respository{App: a}
}

func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Home Page")
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["message"] = "This message comes from template data."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
