package router

import (
	"Gin-Postgres-API/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	//Routers to responde to the API calls from memory
	//To be changed to refer to db
	r := gin.Default()
	r.GET("/people", utils.GetPeople)
	r.GET("/people/id/:id", utils.GetPersonByID)
	r.GET("/people/index/:index", utils.GetPersonByIndex)
	r.GET("/people/guid/:guid", utils.GetPersonByGUID)
	r.GET("/people/isActive/:isActive", utils.GetPersonByIsActive)
	r.GET("/people/balance/:balance", utils.GetPersonByBalance)
	r.GET("/people/age/:age", utils.GetPersonByAge)
	r.GET("/people/eyeColor/:eyeColor", utils.GetPersonByEyeColor)
	//TODO router.GET("/people/name/:name", getPersonByName)
	r.GET("/people/firstname/:firstname", utils.GetPersonByFirstName)
	r.GET("/people/gender/:gender", utils.GetPersonByGender)
	r.GET("/people/company/:company", utils.GetPersonByCompany)
	r.GET("/people/email/:email", utils.GetPersonByEmail)
	r.GET("/people/phoneNumber/:phone", utils.GetPersonByPhoneNumber)
	//TODO router.GET("/people/address/:address", getPersonByAddress)
	r.GET("/people/about/:about", utils.GetPersonByAbout)
	r.GET("/people/registered/:registered", utils.GetPersonByRegistered)
	r.GET("/people/latitude/:latitude", utils.GetPersonByLatitude)
	r.GET("/people/longitude/:longitude", utils.GetPersonByLongitude)
	//TODO router.GET("/people/friends/:friends", getPersonByFriends)
	r.POST("/people", utils.PostPerson)

	return r
}
