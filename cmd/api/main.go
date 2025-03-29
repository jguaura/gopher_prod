package main

import (
	"log"

	"github.com/jguaura/gopher_prod/internal/env"
)

func main() {
	cfg := config {
		address: env.GetString("ADDR", ":8081"),
	}

	app := application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}