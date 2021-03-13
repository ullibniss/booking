package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"github.com/ullibniss/booking/pkg/config"
	"github.com/ullibniss/booking/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}