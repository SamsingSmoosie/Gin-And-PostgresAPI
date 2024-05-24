package main

import (
	"Gin-Postgres-API/internal/repository"
	"Gin-Postgres-API/router"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	_, err := repository.NewPostgresDB("localhost", 5432, "postgres", "TopSecret123!", "postgres")
	if err != nil {
		log.Fatal(err)
	}

	r := router.InitRouter()
	err = r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
