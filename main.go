package main

import (
	"Gin-Postgres-API/internal/repository"
	"Gin-Postgres-API/router"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	db, err := repository.NewPostgresDB(os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreatePeople()
	if nil != err {
		log.Fatal(err)
	}
	err = db.CreateFriends()
	if nil != err {
		log.Fatal(err)
	}
	err = db.CreateMap()
	if nil != err {
		log.Fatal(err)
	}

	r := router.InitRouter(db)
	err = r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
