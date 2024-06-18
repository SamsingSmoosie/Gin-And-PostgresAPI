package main

import (
	"Gin-Postgres-API/internal/repository"
	"Gin-Postgres-API/router"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	db, err := repository.NewPostgresDB("localhost", 5432, "postgres", "TopSecret123!", "postgres")
	if err != nil {
		log.Fatal(err)
	}

	repository.GetJson("data/PersonalData.json")
	err = db.CreatePeople()
	if err != nil {
		log.Fatal(err)
	}
	err = db.CreateFriends()
	if err != nil {
		log.Fatal(err)
	}
	err = db.CreateMap()
	if err != nil {
		log.Fatal(err)
	}

	r := router.InitRouter(db)
	err = r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
