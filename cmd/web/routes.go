package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/purna7788/bookings/pkg/config"
	"github.com/purna7788/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	//mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about",http.HandlerFunc(handlers.Repo.About))

	return mux
}
