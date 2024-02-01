package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

var people []Person

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

func GetPeople(c *gin.Context) {
	c.JSON(http.StatusOK, people)
}

func PostPerson(c *gin.Context) {
	var newPerson Person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	people = append(people, newPerson)
	c.JSON(http.StatusCreated, newPerson)
}

func GetPersonByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range people {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByIndex(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))

	for _, a := range people {
		if a.Index == index {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByGUID(c *gin.Context) {
	guid := c.Param("guid")

	for _, a := range people {
		if a.GUID == guid {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByIsActive(c *gin.Context) {
	isActive, _ := strconv.ParseBool(c.Param("isActive"))
	for _, a := range people {
		if a.IsActive == isActive {
			c.JSON(http.StatusOK, a)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByBalance(c *gin.Context) {
	balance := c.Param("balance")

	for _, a := range people {
		if a.Balance == balance {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByAge(c *gin.Context) {
	age, _ := strconv.Atoi(c.Param("age"))

	for _, a := range people {
		if a.Age == age {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByEyeColor(c *gin.Context) {
	eyeColor := c.Param("eyeColor")

	for _, a := range people {
		if a.EyeColor == eyeColor {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByFirstName(c *gin.Context) {
	firstname := c.Param("firstname")

	for _, a := range people {
		if a.Name.Firstname == firstname {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByGender(c *gin.Context) {
	gender := c.Param("gender")

	for _, a := range people {
		if a.Gender == gender {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByCompany(c *gin.Context) {
	company := c.Param("company")

	for _, a := range people {
		if a.Company == company {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByEmail(c *gin.Context) {
	email := c.Param("isActive")

	for _, a := range people {
		if a.Email == email {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phone")

	for _, a := range people {
		if a.Phone == phoneNumber {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

//TODO Person by address

func GetPersonByAbout(c *gin.Context) {
	about := c.Param("about")

	for _, a := range people {
		if a.About == about {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByRegistered(c *gin.Context) {
	registered := c.Param("registered")

	for _, a := range people {
		if a.Registered == registered {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByLatitude(c *gin.Context) {
	latitude, _ := strconv.ParseFloat(c.Param("latitude"), 64)

	for _, a := range people {
		if a.Latitude == latitude {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func GetPersonByLongitude(c *gin.Context) {
	longitude, _ := strconv.ParseFloat(c.Param("longitude"), 64)

	for _, a := range people {
		if a.Longitude == longitude {
			c.JSON(http.StatusOK, a)

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

//TODO Person by Friends
