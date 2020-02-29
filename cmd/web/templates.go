package main

import (
	"github.com/imranh27/snippetbox/pkg/models"
	"html/template"
	"path/filepath"
)

//define a templateData type to act as the holding structure for any holding data to pass to our HTML templates.
type templateData struct {
	CurrentYear int
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

//Build a template cache
func newTemplateCache(dir string) (map[string]*template.Template, error) {

	//Initialise a new map to act as the cache
	cache := map[string]*template.Template{}

	//Get a slice of all file paths with extension .page.tmpl
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	//Loop through pages one by one
	for _, page := range pages {

		//Extract the name of the page
		name := filepath.Base(page)

		//Parse page template file in to template set.
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		//Add any layout template to the template set.
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		//Add partial templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}

	//return the map
	return cache, nil
}
