package api

import (
	"Gin-Postgres-API/internal/model"
	"Gin-Postgres-API/internal/repository"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var people []model.Person
var repo, _ = repository.NewPostgresDB("localhost", 5432, "postgres", "TopSecret123!", "postgres")
var db = repo.Db

/*// ---------To be Removed once API calls db instead of Memory----------
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

//--------------------------------------------------------------------
*/

// GetPeople Returns all people
func GetPeople(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Println(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var a model.Person
		err := rows.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// PostPerson Can be used to add a person
// Needs to be updated to only allow certain attributes to be imported by user
func PostPerson(c *gin.Context) {
	var newPerson model.Person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	people = append(people, newPerson)
	c.JSON(http.StatusCreated, newPerson)
}

// GetPersonByID Returns one specific person by ID
func GetPersonByID(c *gin.Context) {
	db := repo.Db
	id := c.Param("id")
	row := db.QueryRow("SELECT * FROM people WHERE id = $1", id)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, a)
}

// GetPersonByIndex Returns one specific person by Index
func GetPersonByIndex(c *gin.Context) {
	db := repo.Db
	index, _ := strconv.Atoi(c.Param("index"))

	row, err := db.Query("SELECT * FROM people WHERE index = $1", index)
	if err != nil {
		c.Status(500)
		log.Println(err)
		return
	}

	var a model.Person

	err = row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, a)
}

// GetPersonByGUID Returns every person with a specific GUID
func GetPersonByGUID(c *gin.Context) {
	db := repo.Db
	guid := c.Param("guid")

	row, _ := db.Query("SELECT * FROM people WHERE guid = $1", guid)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByIsActive Returns all people that are active or inactive, depending on input
func GetPersonByIsActive(c *gin.Context) {
	db := repo.Db
	isActive, _ := strconv.ParseBool(c.Param("isActive"))
	row, _ := db.Query("SELECT * FROM people WHERE is_active = $1", isActive)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByBalance Returns all people with a specific balance
func GetPersonByBalance(c *gin.Context) {
	db := repo.Db
	balance := c.Param("balance")

	row, _ := db.Query("SELECT * FROM people WHERE balance = $1", balance)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByAge Returns all people with a specific age
func GetPersonByAge(c *gin.Context) {
	db := repo.Db
	age, _ := strconv.Atoi(c.Param("age"))

	row, _ := db.Query("SELECT * FROM people WHERE age = $1", age)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByEyeColor Returns all people with a specific eyecolor
func GetPersonByEyeColor(c *gin.Context) {
	db := repo.Db
	eyeColor := c.Param("eyeColor")

	row, _ := db.Query("SELECT * FROM people WHERE eye_color = $1", eyeColor)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByFirstName Returns all people with a specific firstname
func GetPersonByFirstName(c *gin.Context) {
	db := repo.Db
	firstname := c.Param("firstname")

	row, _ := db.Query("SELECT * FROM people WHERE name_first = $1", firstname)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLastName Returns all people with a specific lastname
func GetPersonByLastName(c *gin.Context) {
	db := repo.Db
	lastname := c.Param("lastname")

	row, _ := db.Query("SELECT * FROM people WHERE name_first = $1", lastname)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByGender Returns all people with a specific gender
func GetPersonByGender(c *gin.Context) {
	db := repo.Db
	gender := c.Param("gender")

	row, _ := db.Query("SELECT * FROM people WHERE gender = $1", gender)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByCompany Returns all people working for a specific company
func GetPersonByCompany(c *gin.Context) {
	db := repo.Db
	company := c.Param("company")

	row, _ := db.Query("SELECT * FROM people WHERE company = $1", company)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByEmail Returns the person with the specified email address
func GetPersonByEmail(c *gin.Context) {
	db := repo.Db
	email := c.Param("email")

	row, _ := db.Query("SELECT * FROM people WHERE email = $1", email)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByPhoneNumber Returns the person with the specified phone number
func GetPersonByPhoneNumber(c *gin.Context) {
	db := repo.Db
	phoneNumber := c.Param("phone")

	row, _ := db.Query("SELECT * FROM people WHERE phone = $1", phoneNumber)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByHousenumber Returns all people with a specific address
func GetPersonByHousenumber(c *gin.Context) {
	db := repo.Db
	housenumber := c.Param("housenumber")

	row, _ := db.Query("SELECT * FROM people WHERE address_house_number = $1", housenumber)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByStreetname Returns all people with a specific streetname
func GetPersonByStreetname(c *gin.Context) {
	db := repo.Db
	streetname := c.Param("streetname")

	row, _ := db.Query("SELECT * FROM people WHERE address_street = $1", streetname)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByZipcode Returns all people with a specific address
func GetPersonByZipcode(c *gin.Context) {
	db := repo.Db
	zipcode := c.Param("zipcode")

	row, _ := db.Query("SELECT * FROM people WHERE address_zip_code = $1", zipcode)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByCity Returns all people with a specific city
func GetPersonByCity(c *gin.Context) {
	db := repo.Db
	city := c.Param("city")

	row, _ := db.Query("SELECT * FROM people WHERE address_city = $1", city)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByState Returns all people with a specific address
func GetPersonByState(c *gin.Context) {
	db := repo.Db
	state := c.Param("state")

	row, _ := db.Query("SELECT * FROM people WHERE address_state = $1", state)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByAbout Returns all people with a specific about
func GetPersonByAbout(c *gin.Context) {
	db := repo.Db
	about := c.Param("about")

	row, _ := db.Query("SELECT * FROM people WHERE about = $1", about)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByRegistered Returns all people with a specific registration
func GetPersonByRegistered(c *gin.Context) {
	db := repo.Db
	registered := c.Param("registered")

	row, _ := db.Query("SELECT * FROM people WHERE registered = $1", registered)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLatitude Returns all people with a specific latitude
func GetPersonByLatitude(c *gin.Context) {
	db := repo.Db
	latitude, _ := strconv.ParseFloat(c.Param("latitude"), 64)

	row, _ := db.Query("SELECT * FROM people WHERE latitude = $1", latitude)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByLongitude Returns all people with a specific longitude
func GetPersonByLongitude(c *gin.Context) {
	db := repo.Db
	longitude, _ := strconv.ParseFloat(c.Param("longitude"), 64)

	row, _ := db.Query("SELECT * FROM people WHERE longitude = $1", longitude)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
		return
	}

	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			log.Println(err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
}

//TODO Person by Friends
