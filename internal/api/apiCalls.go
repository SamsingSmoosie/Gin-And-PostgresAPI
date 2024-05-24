package api

import (
	"Gin-Postgres-API/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//var people []model.Person

/*   ---------To be Removed once API calls db instead of Memory----------
func GetJson(filepath string) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &people)
	if err != nil {
		log.Fatal(err)
	}
}

*/ //--------------------------------------------------------------------

// GetPeople Returns all people
func GetPeople(c *gin.Context) {
	people, err := repository.Postgres.GetPeople()
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
func GetPersonByID(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByID()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByIndex Returns one specific person by Index
func GetPersonByIndex(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByIndex()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByGUID Returns every person with a specific GUID
func GetPersonByGUID(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByGUID()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByIsActive Returns all people that are active or inactive, depending on input
func GetPersonByIsActive(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByIsActive()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByBalance Returns all people with a specific balance
func GetPersonByBalance(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByBalance()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByAge Returns all people with a specific age
func GetPersonByAge(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByAge()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByEyeColor Returns all people with a specific eyecolor
func GetPersonByEyeColor(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByEyeColor()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByFirstName Returns all people with a specific firstname
func GetPersonByFirstName(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByFirstName()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLastName Returns all people with a specific lastname
func GetPersonByLastName(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByLastName()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByGender Returns all people with a specific gender
func GetPersonByGender(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByGender()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByCompany Returns all people working for a specific company
func GetPersonByCompany(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByCompany()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByEmail Returns the person with the specified email address
func GetPersonByEmail(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByEmail()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByPhoneNumber Returns the person with the specified phone number
func GetPersonByPhoneNumber(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByPhoneNumber()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByHousenumber Returns all people with a specific address
func GetPersonByHousenumber(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByHouseNumber()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByStreetname Returns all people with a specific streetname
func GetPersonByStreetname(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByStreetName()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByZipcode Returns all people with a specific address
func GetPersonByZipcode(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByZipCode()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByCity Returns all people with a specific city
func GetPersonByCity(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByCity()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByState Returns all people with a specific address
func GetPersonByState(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByState()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByAbout Returns all people with a specific about
func GetPersonByAbout(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByAbout()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByRegistered Returns all people with a specific registration
func GetPersonByRegistered(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByRegistered()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLatitude Returns all people with a specific latitude
func GetPersonByLatitude(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByLatitude()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLongitude Returns all people with a specific longitude
func GetPersonByLongitude(c *gin.Context) {
	people, err := repository.Postgres.GetPersonByLongitude()

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, people)
}

//TODO Person by Friends
