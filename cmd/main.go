package main

import (
	"database/sql"
	"log"

	"github.com/HerbertSanches/ecommerce-go/cmd/api"
	"github.com/HerbertSanches/ecommerce-go/config"
	"github.com/HerbertSanches/ecommerce-go/db"
	_ "github.com/lib/pq"
)

func main() {
	db, err := db.NewPostgreStorage(config.Envs)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)	
	}

	log.Println("DB: Successfully connected!")
}