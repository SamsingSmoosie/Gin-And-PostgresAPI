package repository

import (
	"Gin-Postgres-API/internal/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgresDB(host string, port int, user, password, dbname string) (*Postgres, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	openedDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = openedDB.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Database responds to ping")

	return &Postgres{db: openedDB}, err
}

// CreatePeople Check if table "people" exists. If not it will create the table and populate it
func (p *Postgres) createPeople(people []model.Person) error {
	_, tableCheck := p.db.Query("select * from people;")
	if tableCheck == nil {
		log.Println("Table people already exists")
		return nil
	}

	_, err := p.db.Exec("CREATE TABLE people(id varchar primary key, index int, guid varchar, is_active bool, balance varchar, picture varchar, age int, eye_color varchar, name_first varchar, name_last varchar, gender varchar, company varchar, email varchar, phone varchar, address_house_number int, address_street varchar, address_city varchar, address_state varchar, address_zip_code int, about varchar, registered varchar, latitude float8, longitude float8)")
	if err != nil {
		return err
	}
	for _, person := range people {
		p.insertPerson(person)
	}
	log.Println("Table people created and populated")
	return nil
}

// CreateFriends Check if table "friend" exists. If not it will create the table and populate it
func (p *Postgres) createFriends(people []model.Person) error {
	_, tableCheck := p.db.Query("select * from friends;")
	if tableCheck == nil {
		log.Println("Table friends already exists")
		return nil
	}
	_, err := p.db.Exec("CREATE TABLE friends(id varchar primary key, name_first varchar, name_last varchar)")
	if err != nil {
		return err
	}
	for _, person := range people {
		p.insertFriends(person)
	}
	log.Println("Table friend created and populated")
	return nil
}

// CreateMap Check if table "person_friend_map" exists. If not it will create the table and populate it
func (p *Postgres) createMap(people []model.Person) error {
	_, tableCheck := p.db.Query("select * from person_friend_map;")
	if tableCheck == nil {
		log.Println("Table person already exists")
		return nil
	}
	_, err := p.db.Exec(`CREATE TABLE person_friend_map(person_id varchar, friend_id varchar, CONSTRAINT person_id_FK FOREIGN KEY (person_id) REFERENCES people (id), CONSTRAINT friend_id_FK FOREIGN KEY (friend_id) REFERENCES friends (id))`)
	if err != nil {
		return err
	}
	for _, person := range people {
		p.insertMap(person)
	}
	log.Println("Table person_friend_map created and populated")
	return nil
}

func (p *Postgres) insertPerson(person model.Person) {
	insertStatement := `
		INSERT INTO people (
			id, index, guid, is_active, balance, picture, age, eye_color,
			name_first, name_last, gender, company, email, phone,
			address_house_number, address_street, address_city,
			address_state, address_zip_code, about, registered,
			latitude, longitude
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
	`
	_, err := p.db.Exec(insertStatement,
		person.ID, person.Index, person.GUID, person.IsActive, person.Balance, person.Picture, person.Age, person.EyeColor,
		person.Name.Firstname, person.Name.Lastname, person.Gender, person.Company, person.Email, person.Phone,
		person.Address.HouseNumber, person.Address.Street, person.Address.City,
		person.Address.State, person.Address.ZipCode, person.About, person.Registered,
		person.Latitude, person.Longitude,
	)

	if err != nil {
		log.Printf("Error inserting person: %v\n", err)
	}
}

func (p *Postgres) insertFriends(person model.Person) {

	insertStatement := `INSERT INTO friends (id, name_first, name_last) VALUES ($1, $2, $3)`

	for _, friend := range person.Friends {
		_, err := p.db.Exec(insertStatement, friend.ID, friend.Name.Firstname, friend.Name.Lastname)
		if err != nil {
			log.Printf("Error inserting friends: %v\n", err)
		}
	}
}

func (p *Postgres) insertMap(person model.Person) {

	insertStatement := `INSERT INTO person_friend_map (person_id, friend_id) VALUES ($1, $2)`

	for _, friend := range person.Friends {
		_, err := p.db.Exec(insertStatement, person.ID, friend.ID)
		if err != nil {
			log.Printf("Error inserting person: %s", err)
		}
	}
}

func (p *Postgres) GetPeople(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	rows, err := p.db.Query("SELECT * FROM people")
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var a model.Person
		err = rows.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return nil, fmt.Errorf("row scan: %w", err)
		}
		people = append(people, a)
	}
	c.JSON(http.StatusOK, people)
	return people, nil
}

func (p *Postgres) GetPersonByID(c *gin.Context) (model.Person, error) {
	db := p.db
	id := c.Param("id")
	row := db.QueryRow("SELECT * FROM people WHERE id = $1", id)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
	}

	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, a)
	return a, nil
}

func (p *Postgres) GetPersonByIndex(c *gin.Context) (model.Person, error) {
	db := p.db
	index, _ := strconv.Atoi(c.Param("index"))

	row, err := db.Query("SELECT * FROM people WHERE index = $1", index)
	if err != nil {
		c.Status(500)
		log.Println(err)
	}

	var a model.Person

	err = row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, a)
	return a, nil
}

func (p *Postgres) GetPersonByGUID(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	guid := c.Param("guid")

	row, err := db.Query("SELECT * FROM people WHERE guid = $1", guid)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByIsActive(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	isActive, _ := strconv.ParseBool(c.Param("isActive"))

	row, err := db.Query("SELECT * FROM people WHERE is_active = $1", isActive)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByBalance(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	balance, _ := strconv.ParseBool(c.Param("balance"))

	row, err := db.Query("SELECT * FROM people WHERE balance = $1", balance)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByAge(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	age, _ := strconv.ParseBool(c.Param("age"))

	row, err := db.Query("SELECT * FROM people WHERE age = $1", age)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByEyeColor(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	eyeColor, _ := strconv.ParseBool(c.Param("eyeColor"))

	row, err := db.Query("SELECT * FROM people WHERE eye_color = $1", eyeColor)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByFirstName(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	firstname, _ := strconv.ParseBool(c.Param("firstname"))

	row, err := db.Query("SELECT * FROM people WHERE name_first = $1", firstname)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByLastName(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	lastname, _ := strconv.ParseBool(c.Param("lastname"))

	row, err := db.Query("SELECT * FROM people WHERE name_last = $1", lastname)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByGender(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	gender, _ := strconv.ParseBool(c.Param("gender"))

	row, err := db.Query("SELECT * FROM people WHERE gender = $1", gender)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByCompany(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	company, _ := strconv.ParseBool(c.Param("company"))

	row, err := db.Query("SELECT * FROM people WHERE company = $1", company)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByEmail(c *gin.Context) (model.Person, error) {
	db := p.db
	email := c.Param("email")
	row := db.QueryRow("SELECT * FROM people WHERE email = $1", email)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
	}

	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, a)
	return a, nil
}

func (p *Postgres) GetPersonByPhoneNumber(c *gin.Context) (model.Person, error) {
	db := p.db
	phoneNumber := c.Param("phone")
	row := db.QueryRow("SELECT * FROM people WHERE phone = $1", phoneNumber)
	if row.Err() != nil {
		c.Status(500)
		log.Println(row.Err())
	}

	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, a)
	return a, nil
}

func (p *Postgres) GetPersonByHouseNumber(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	houseNumber, _ := strconv.ParseBool(c.Param("housenumber"))

	row, err := db.Query("SELECT * FROM people WHERE address_house_number = $1", houseNumber)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByStreetName(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	streetName, _ := strconv.ParseBool(c.Param("streetname"))

	row, err := db.Query("SELECT * FROM people WHERE address_street = $1", streetName)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByZipCode(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	zipCode, _ := strconv.ParseBool(c.Param("zipcode"))

	row, err := db.Query("SELECT * FROM people WHERE address_zip_code = $1", zipCode)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByCity(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	city, _ := strconv.ParseBool(c.Param("city"))

	row, err := db.Query("SELECT * FROM people WHERE address_city = $1", city)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByState(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	state, _ := strconv.ParseBool(c.Param("state"))

	row, err := db.Query("SELECT * FROM people WHERE address_state = $1", state)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByAbout(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	about, _ := strconv.ParseBool(c.Param("about"))

	row, err := db.Query("SELECT * FROM people WHERE about = $1", about)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByRegistered(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	registered, _ := strconv.ParseBool(c.Param("registered"))

	row, err := db.Query("SELECT * FROM people WHERE registered = $1", registered)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByLatitude(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	latitude, _ := strconv.ParseBool(c.Param("latitude"))

	row, err := db.Query("SELECT * FROM people WHERE latitude = $1", latitude)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}

func (p *Postgres) GetPersonByLongitude(c *gin.Context) ([]model.Person, error) {
	var people []model.Person
	db := p.db
	longitude, _ := strconv.ParseBool(c.Param("longitude"))

	row, err := db.Query("SELECT * FROM people WHERE longitude = $1", longitude)
	if err != nil {
		c.Status(500)
		log.Println(err)
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
	return people, nil
}
