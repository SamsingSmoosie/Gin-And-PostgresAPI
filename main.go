package main

import (
	"Gin-Postgres-API/internal/api"
	"Gin-Postgres-API/internal/repository"
	"Gin-Postgres-API/router"
	_ "github.com/lib/pq"
	"log"
)

var JsonFile = "./data/PersonalData.json"

func main() {

	db, err := repository.NewPostgresDB("localhost", 5432, "postgres", "TopSecret123!", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	api.NewAPI(db)

	r := router.InitRouter()
	err = r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
