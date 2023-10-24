package main

import (
	"lec_1/pkg/config"
	"lec_1/pkg/handlers"
	"lec_1/pkg/render"
	"lec_1/pkg/routes"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCash()
	if err != nil {
		log.Fatal(err.Error())
	}
	app.TamplateCashe = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	url := "127.0.0.1:8080"
	srv := http.Server{
		Addr:    url,
		Handler: routes.MyRoutes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
