package router

import (
	"Gin-Postgres-API/internal/api"
	"Gin-Postgres-API/internal/repository"
	"github.com/gin-gonic/gin"
)

func InitRouter(db *repository.Postgres) *gin.Engine {

	//Routers to responde to the API calls from memory
	//To be changed to refer to db

	handlers := &api.Handlers{DB: db}

	r := gin.Default()
	r.GET("/people", handlers.GetPeople)
	r.GET("/people/id/:id", handlers.GetPersonByID)
	r.GET("/people/index/:index", handlers.GetPersonByIndex)
	r.GET("/people/guid/:guid", handlers.GetPersonByGUID)
	r.GET("/people/isActive/:isActive", handlers.GetPersonByIsActive)
	r.GET("/people/balance/:balance", handlers.GetPersonByBalance)
	r.GET("/people/age/:age", handlers.GetPersonByAge)
	r.GET("/people/eyeColor/:eyeColor", handlers.GetPersonByEyeColor)
	r.GET("/people/lastname/:lastname", handlers.GetPersonByLastName)
	r.GET("/people/firstname/:firstname", handlers.GetPersonByFirstName)
	r.GET("/people/gender/:gender", handlers.GetPersonByGender)
	r.GET("/people/company/:company", handlers.GetPersonByCompany)
	r.GET("/people/email/:email", handlers.GetPersonByEmail)
	r.GET("/people/phoneNumber/:phone", handlers.GetPersonByPhoneNumber)
	r.GET("/people/housenumber/:housenumber", handlers.GetPersonByHousenumber)
	r.GET("/people/streetname/:streetname", handlers.GetPersonByStreetname)
	r.GET("/people/zipcode/:zipcode", handlers.GetPersonByZipcode)
	r.GET("/people/city/:city", handlers.GetPersonByCity)
	r.GET("/people/state/:state", handlers.GetPersonByState)
	r.GET("/people/about/:about", handlers.GetPersonByAbout)
	r.GET("/people/registered/:registered", handlers.GetPersonByRegistered)
	r.GET("/people/latitude/:latitude", handlers.GetPersonByLatitude)
	r.GET("/people/longitude/:longitude", handlers.GetPersonByLongitude)
	//TODO router.GET("/people/friends/:friends", getPersonByFriends)

	//r.POST("/people", api.PostPerson)

	return r
}
