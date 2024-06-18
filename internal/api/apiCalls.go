package api

import (
	"Gin-Postgres-API/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Handlers struct {
	DB *repository.Postgres
}

// GetPeople Returns all people
func (h *Handlers) GetPeople(c *gin.Context) {
	people, err := h.DB.GetPeople()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// PostPerson Can be used to add a person
// Needs to be updated to only allow certain attributes to be imported by user
/*func PostPerson(c *gin.Context) {
	var newPerson model.Person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	people = append(people, newPerson)
	c.JSON(http.StatusCreated, newPerson)
}*/

// GetPersonByID Returns one specific person by ID
func (h *Handlers) GetPersonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	people, err := h.DB.GetPersonByID(id)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByIndex Returns one specific person by Index
func (h *Handlers) GetPersonByIndex(c *gin.Context) {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		log.Println(err)
		return
	}
	people, err := h.DB.GetPersonByIndex(index)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByGUID Returns every person with a specific GUID
func (h *Handlers) GetPersonByGUID(c *gin.Context) {
	guid := c.Param("guid")
	people, err := h.DB.GetPersonByGUID(guid)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByIsActive Returns all people that are active or inactive, depending on input
func (h *Handlers) GetPersonByIsActive(c *gin.Context) {
	isActive, err := strconv.ParseBool(c.Param("is_active"))
	if err != nil {
		log.Println(err)
		return
	}
	people, err := h.DB.GetPersonByIsActive(isActive)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByBalance Returns all people with a specific balance
func (h *Handlers) GetPersonByBalance(c *gin.Context) {
	balance := c.Param("balance")
	people, err := h.DB.GetPersonByBalance(balance)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByAge Returns all people with a specific age
func (h *Handlers) GetPersonByAge(c *gin.Context) {
	age, err := strconv.Atoi(c.Param("age"))
	if err != nil {
		log.Println(err)
		return
	}
	people, err := h.DB.GetPersonByAge(age)

	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByEyeColor Returns all people with a specific eyecolor
func (h *Handlers) GetPersonByEyeColor(c *gin.Context) {
	eyeColor := c.Param("eye_color")
	people, err := h.DB.GetPersonByEyeColor(eyeColor)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByFirstName Returns all people with a specific firstname
func (h *Handlers) GetPersonByFirstName(c *gin.Context) {
	firstName := c.Param("first_name")
	people, err := h.DB.GetPersonByFirstName(firstName)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLastName Returns all people with a specific lastname
func (h *Handlers) GetPersonByLastName(c *gin.Context) {
	lastName := c.Param("last_name")
	people, err := h.DB.GetPersonByLastName(lastName)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByGender Returns all people with a specific gender
func (h *Handlers) GetPersonByGender(c *gin.Context) {
	gender := c.Param("gender")
	people, err := h.DB.GetPersonByGender(gender)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByCompany Returns all people working for a specific company
func (h *Handlers) GetPersonByCompany(c *gin.Context) {
	company := c.Param("company")
	people, err := h.DB.GetPersonByCompany(company)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByEmail Returns the person with the specified email address
func (h *Handlers) GetPersonByEmail(c *gin.Context) {
	email := c.Param("email")
	people, err := h.DB.GetPersonByEmail(email)

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByPhoneNumber Returns the person with the specified phone number
func (h *Handlers) GetPersonByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phone_number")
	people, err := h.DB.GetPersonByPhoneNumber(phoneNumber)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByHousenumber Returns all people with a specific address
func (h *Handlers) GetPersonByHousenumber(c *gin.Context) {
	housenumber, err := strconv.Atoi(c.Param("housenumber"))
	if err != nil {
		log.Println(err)
		return
	}
	people, err := h.DB.GetPersonByHouseNumber(housenumber)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByStreetname Returns all people with a specific streetname
func (h *Handlers) GetPersonByStreetname(c *gin.Context) {
	streetname := c.Param("streetname")
	people, err := h.DB.GetPersonByStreetName(streetname)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByZipcode Returns all people with a specific address
func (h *Handlers) GetPersonByZipcode(c *gin.Context) {
	zipcode, err := strconv.Atoi(c.Param("zipcode"))
	if err != nil {
		log.Println(err)
		return
	}
	people, err := h.DB.GetPersonByZipCode(zipcode)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByCity Returns all people with a specific city
func (h *Handlers) GetPersonByCity(c *gin.Context) {
	city := c.Param("city")
	people, err := h.DB.GetPersonByCity(city)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByState Returns all people with a specific address
func (h *Handlers) GetPersonByState(c *gin.Context) {
	state := c.Param("state")
	people, err := h.DB.GetPersonByState(state)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByAbout Returns all people with a specific about
func (h *Handlers) GetPersonByAbout(c *gin.Context) {
	about := c.Param("about")
	people, err := h.DB.GetPersonByAbout(about)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByRegistered Returns all people with a specific registration
func (h *Handlers) GetPersonByRegistered(c *gin.Context) {
	registered, err := strconv.ParseBool(c.Param("registered"))
	if err != nil {
		log.Println(err)
		return
	}
	people, err := h.DB.GetPersonByRegistered(registered)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLatitude Returns all people with a specific latitude
func (h *Handlers) GetPersonByLatitude(c *gin.Context) {
	latitude := c.Param("latitude")
	people, err := h.DB.GetPersonByLatitude(latitude)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLongitude Returns all people with a specific longitude
func (h *Handlers) GetPersonByLongitude(c *gin.Context) {
	longitude := c.Param("longitude")
	people, err := h.DB.GetPersonByLongitude(longitude)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, people)
}

//TODO Person by Friends
