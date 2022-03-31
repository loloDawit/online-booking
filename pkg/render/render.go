package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func Template(w http.ResponseWriter, tmpl string) {

	tc, err := CreatTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	_ = t.Execute(buffer, nil)

	_, err = buffer.WriteTo(w)

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
