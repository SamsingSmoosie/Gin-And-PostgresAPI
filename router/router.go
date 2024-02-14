package router

import (
	"Gin-Postgres-API/internal/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	//Routers to responde to the API calls from memory
	//To be changed to refer to db
	r := gin.Default()
	r.GET("/people", api.GetPeople)
	r.GET("/people/id/:id", api.GetPersonByID)
	r.GET("/people/index/:index", api.GetPersonByIndex)
	r.GET("/people/guid/:guid", api.GetPersonByGUID)
	r.GET("/people/isActive/:isActive", api.GetPersonByIsActive)
	r.GET("/people/balance/:balance", api.GetPersonByBalance)
	r.GET("/people/age/:age", api.GetPersonByAge)
	r.GET("/people/eyeColor/:eyeColor", api.GetPersonByEyeColor)
	r.GET("/people/lastname/:lastname", api.GetPersonByLastName)
	r.GET("/people/firstname/:firstname", api.GetPersonByFirstName)
	r.GET("/people/gender/:gender", api.GetPersonByGender)
	r.GET("/people/company/:company", api.GetPersonByCompany)
	r.GET("/people/email/:email", api.GetPersonByEmail)
	r.GET("/people/phoneNumber/:phone", api.GetPersonByPhoneNumber)
	r.GET("/people/housenumber/:housenumber", api.GetPersonByHousenumber)
	r.GET("/people/streetname/:streetname", api.GetPersonByStreetname)
	r.GET("/people/zipcode/:zipcode", api.GetPersonByZipcode)
	r.GET("/people/city/:city", api.GetPersonByCity)
	r.GET("/people/state/:state", api.GetPersonByState)
	r.GET("/people/about/:about", api.GetPersonByAbout)
	r.GET("/people/registered/:registered", api.GetPersonByRegistered)
	r.GET("/people/latitude/:latitude", api.GetPersonByLatitude)
	r.GET("/people/longitude/:longitude", api.GetPersonByLongitude)
	//TODO router.GET("/people/friends/:friends", getPersonByFriends)

	r.POST("/people", api.PostPerson)

	return r
}
