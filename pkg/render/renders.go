package render

import (
	"bytes"
	"html/template"
	"lec_1/pkg/config"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, temp string) {
	tc, err := CreateTemplateCash()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[temp]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCash() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return mycache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}

		maches, err := filepath.Glob("./templates/*.layout.templ")

		if err != nil {
			return mycache, err
		}

		if len(maches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.templ")
			if err != nil {
				return mycache, err
			}
		}
		mycache[name] = ts
	}
	return mycache, nil
}
