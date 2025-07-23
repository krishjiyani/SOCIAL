package main

import (
	"log"
	"krishjiyani/SOCIAL/internal/env"
	"krishjiyani/SOCIAL/internal/store"
	"krishjiyani/SOCIAL/internal/db"
)

func main() {
	addr := env.GetString("DB_ADDR","postgres://postgres:Krish.jiyani%401@localhost:5432/socialnetworkx_ai?sslmode=disable" )
	conn, err := db.New(addr, 30, 30,"15m")
    if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)

}