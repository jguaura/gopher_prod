package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jguaura/gopher_prod/internal/store"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	address string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route(("/v1"), func(r chi.Router) {
		r.Get("/health", app.healtCheckHandler)
	})
		
	return r
}

func (app *application) run(mux http.Handler) error {
	
	srv := &http.Server{
		Addr: app.config.address,
		Handler: mux,
		WriteTimeout: time.Second * 15,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started on %s", app.config.address)

	return srv.ListenAndServe()
}