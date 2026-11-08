package main

import (
	"html/template"
	"path/filepath"

	"github.com/ctailee/snippetbox/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./assets/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)

		ts, err := template.ParseFiles("assets/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("assets/html/components/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[filename] = ts
	}

	return cache, nil
}
