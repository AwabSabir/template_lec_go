package main

import (
	"fmt"
	"lec_1/pkg/config"
	"lec_1/pkg/handler"
	"lec_1/pkg/render"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hi there is this is my first lec for advance web")
	var app config.AppConfig
	tc, err := render.CreateTemplateCash()
	if err != nil {
		log.Fatal(err.Error())
	}
	app.TamplateCashe = tc
	app.UseCache = false
	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	render.NewTemplate(&app)
	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)
	url := "127.0.0.1:8080"
	_ = http.ListenAndServe(url, nil)
}
