package main

import (
	"Gin-Postgres-API/dbpop"
	"Gin-Postgres-API/internal/repository"
	"Gin-Postgres-API/router"
	_ "github.com/lib/pq"
	"log"
)

var JsonFile = "./data/PersonalData.json"

func main() {

	repo, err := repository.NewPostgresDB("localhost", 5432, "postgres", "TopSecret123!", "postgres")
	if err != nil {
		log.Fatal(err)
	}

	db := repo.Db

	err = dbpop.Create(JsonFile, db)
	if err != nil {
		log.Fatal(err)
	}

	r := router.InitRouter()
	err = r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
