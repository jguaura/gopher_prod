package main

import (
	"log"

	"github.com/jguaura/gopher_prod/internal/env"
	"github.com/jguaura/gopher_prod/internal/store"
)

func main() {
	cfg := config {
		address: env.GetString("ADDR", ":8081"),
	}
	
	store := store.NewStorage(nil)
	
	app := application{
		config: cfg,
		store: store,
	}
	
	mux := app.mount()

	log.Fatal(app.run(mux))
}