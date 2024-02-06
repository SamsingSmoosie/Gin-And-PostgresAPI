package router

import (
	"Gin-Postgres-API/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	//Routers to responde to the API calls from memory
	//To be changed to refer to db
	router := gin.Default()
	router.GET("/people", utils.GetPeople)
	router.GET("/people/id/:id", utils.GetPersonByID)
	router.GET("/people/index/:index", utils.GetPersonByIndex)
	router.GET("/people/guid/:guid", utils.GetPersonByGUID)
	router.GET("/people/isActive/:isActive", utils.GetPersonByIsActive)
	router.GET("/people/balance/:balance", utils.GetPersonByBalance)
	router.GET("/people/age/:age", utils.GetPersonByAge)
	router.GET("/people/eyeColor/:eyeColor", utils.GetPersonByEyeColor)
	//TODO router.GET("/people/name/:name", getPersonByName)
	router.GET("/people/firstname/:firstname", utils.GetPersonByFirstName)
	router.GET("/people/gender/:gender", utils.GetPersonByGender)
	router.GET("/people/company/:company", utils.GetPersonByCompany)
	router.GET("/people/email/:email", utils.GetPersonByEmail)
	router.GET("/people/phoneNumber/:phone", utils.GetPersonByPhoneNumber)
	//TODO router.GET("/people/address/:address", getPersonByAddress)
	router.GET("/people/about/:about", utils.GetPersonByAbout)
	router.GET("/people/registered/:registered", utils.GetPersonByRegistered)
	router.GET("/people/latitude/:latitude", utils.GetPersonByLatitude)
	router.GET("/people/longitude/:longitude", utils.GetPersonByLongitude)
	//TODO router.GET("/people/friends/:friends", getPersonByFriends)
	router.POST("/people", utils.PostPerson)

	return router
}
