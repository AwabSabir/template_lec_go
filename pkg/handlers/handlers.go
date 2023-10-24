package handlers

import (
	"lec_1/pkg/config"
	"lec_1/pkg/models"
	"lec_1/pkg/render"
	"net/http"
)

//send data to template

var Repo *Repostory

type Repostory struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repostory {
	return &Repostory{
		App: a,
	}
}

func NewHandler(r *Repostory) {
	Repo = r
}

func (m *Repostory) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repostory) About(w http.ResponseWriter, r *http.Request) {
	//logic
	StringMap := make(map[string]string)
	StringMap["test"] = "hello again"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StrinMap: StringMap,
	})
}
