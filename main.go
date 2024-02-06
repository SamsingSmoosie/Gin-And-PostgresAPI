package main

import (
	"Gin-Postgres-API/dbpop"
	"Gin-Postgres-API/router"
	"Gin-Postgres-API/utils"
	_ "github.com/lib/pq"
)

var JsonFile = "./data/PersonalData.json"

func main() {

	utils.GetJson(JsonFile) //To be removed once the api calls the db not the memory
	dbpop.Create(JsonFile)

	router := router.InitRouter()
	router.Run("localhost:8080")

}
