package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"gitlab.nordstrom.com/online-booking/pkg/config"
	"gitlab.nordstrom.com/online-booking/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// sets the config for the template pkg
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData to make data sharing easy
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func Template(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	// get the template cache from the app config
	// tc, err := CreatTemplateCache() ... this works but it loads everytime we load a page and its not efficent
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache // read info from the cache otherwise rebuild
	} else {
		tc, _ = CreatTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buffer := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)
	_ = t.Execute(buffer, templateData)

	_, err := buffer.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreatTemplateCache() (map[string]*template.Template, error) {
	// look up by name of the template
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return cache, err
			}
			cache[name] = ts
		}

	}
	return cache, err
}
