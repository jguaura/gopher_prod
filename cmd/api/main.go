package main

import (
	"log"

	"github.com/jguaura/gopher_prod/internal/db"
	"github.com/jguaura/gopher_prod/internal/env"
	"github.com/jguaura/gopher_prod/internal/store"
)

func main() {
	cfg := config {
		address: env.GetString("ADDR", ":8081"),
		db: dbConfig{
			address: env.GetString("DB_ADDRESS", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 20),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 20),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}
	
	db, err := db.New(
		cfg.db.address,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("db connection established")
	store := store.NewStorage(db)
	
	app := application{
		config: cfg,
		store: store,
	}
	
	mux := app.mount()

	log.Fatal(app.run(mux))
}