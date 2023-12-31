package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"web-templates/pkg/config"
	"web-templates/pkg/models"
)

var app *config.AppConfig

// Getting app instance of struct from main package to render packege.
// Cant make app in main package public and use it here in RenderTemplate func, as it will give cyclic import errors.
func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	// if UseCache is true, then Cache is built once, and used for all page requests.
	// Cache is not rebuilt again on every page refresh to check realtime changes to html pages.
	if app.UseCache {
		parsedTemplate := app.TemplateCache[tmpl]
		err := parsedTemplate.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		templateMap, _ := CreateTemplateCache()
		parsedTemplate := templateMap[tmpl]
		err := parsedTemplate.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}

	layoutfile, err := filepath.Glob("../../templates/*.layout.tmpl")
	if err != nil {
		fmt.Println(err)
		return tc, err
	}
	fmt.Println("Layout files:", layoutfile)
	pagefiles, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		fmt.Println(err)
		return tc, err
	}
	fmt.Println("Page files:", pagefiles)

	for _, page := range pagefiles {
		pagekey := filepath.Base(page)
		fmt.Println("Templating for", pagekey)
		if len(layoutfile) > 0 {
			intermediateTemplate, err := template.ParseFiles(page)
			if err != nil {
				fmt.Println(err)
				return tc, err
			}
			tc[pagekey], err = intermediateTemplate.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				fmt.Println(err)
				return tc, err
			}
		} else {
			tc[pagekey], err = template.ParseFiles(page)
			if err != nil {
				fmt.Println(err)
				return tc, err
			}
		}
	}

	return tc, nil
}
