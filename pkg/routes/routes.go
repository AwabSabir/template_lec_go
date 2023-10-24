package routes

import (
	"github.com/bmizerany/pat"
	"lec_1/pkg/config"
	"lec_1/pkg/handlers"
	"net/http"
)

func MyRoutes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
