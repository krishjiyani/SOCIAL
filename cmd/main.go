package main

import (
	"krishjiyani/SOCIAL/cmd/api"
	"krishjiyani/SOCIAL/internal/db"
	"krishjiyani/SOCIAL/internal/env"
	"krishjiyani/SOCIAL/internal/store"
	"log"
)

const version = "0.0.1"

func main() {
	cfg := api.Config{
		Addr: env.GetString("ADDR", ":8080"),
		Db: api.DbConfig{
			Addr:         env.GetString("DB_ADDR", "postgres://postgres:Krish.jiyani%401@localhost:5432/socialnetworkx_ai?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		Env: env.GetString("ENV", "development"),
	}

	database, err := db.New(
		cfg.Db.Addr,
		cfg.Db.MaxOpenConns,
		cfg.Db.MaxIdleConns,
		cfg.Db.MaxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}
	defer database.Close()
	log.Println("Database connection pool initialized")

	store := store.NewStorage(database)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
