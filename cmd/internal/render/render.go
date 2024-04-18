package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/purna7788/bookings/cmd/internal/config"
	"github.com/purna7788/bookings/cmd/internal/models"
)

var tc *config.AppConfig

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td models.TemplateData) error {

	var cache map[string]*template.Template
	var err error
	if tc.UseCache {
		cache = tc.TemplateCache
	} else {
		cache, err = CreateCache()
		if err != nil {
			log.Fatal("not able to create template cache")
		}
	}

	tl, ok := cache[tmpl]

	if !ok {
		log.Fatal("not able to create template cache")
	}

	td = dafaultData(td, r)
	var buf = new(bytes.Buffer)
	err = tl.Execute(buf, td)
	if err != nil {
		return err
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}

func CreateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	allPages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil {
		return cache, err
	}

	basePagePath, err := filepath.Glob("./templates/*layout.tmpl")
	if err != nil {
		return cache, err
	}

	if len(basePagePath) == 0 {
		return cache, err
	}

	for _, page := range allPages {
		name := filepath.Base(page)
		template, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		template, err = template.ParseFiles(basePagePath[0])
		if err != nil {
			return cache, err
		}
		cache[name] = template
	}
	return cache, nil
}

func NewTemplate(a *config.AppConfig) {
	tc = a
}

func dafaultData(td models.TemplateData, r *http.Request) models.TemplateData {
	csrfToken := nosurf.Token(r)
	td.CSRFToken = csrfToken

	return td
}
