package repository

import (
	"Gin-Postgres-API/internal/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Postgres struct {
	db *sql.DB
}

var people []model.Person

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

func checkJson(filepath string) error {
	if len(people) != 0 {
		log.Println("Json has already been parsed")
		return nil
	} else {
		file, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(file, &people)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreatePeople Check if table "people" exists. If not it will create the table and populate it
func (p *Postgres) CreatePeople() error {
	_, tableCheck := p.db.Query("select * from people;")
	if tableCheck == nil {
		log.Println("Table people already exists")
		return nil
	}
	err := checkJson("data/PersonalData.json")
	if err != nil {
		return err
	}
	_, err = p.db.Exec("CREATE TABLE people(id varchar primary key, index int, guid varchar, is_active bool, balance varchar, picture varchar, age int, eye_color varchar, name_first varchar, name_last varchar, gender varchar, company varchar, email varchar, phone varchar, address_house_number int, address_street varchar, address_city varchar, address_state varchar, address_zip_code int, about varchar, registered varchar, latitude float8, longitude float8)")
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
func (p *Postgres) CreateFriends() error {
	_, tableCheck := p.db.Query("select * from friends;")
	if tableCheck == nil {
		log.Println("Table friends already exists")
		return nil
	}
	err := checkJson("data/PersonalData.json")
	if err != nil {
		return err
	}
	_, err = p.db.Exec("CREATE TABLE friends(id varchar primary key, name_first varchar, name_last varchar)")
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
func (p *Postgres) CreateMap() error {
	_, tableCheck := p.db.Query("select * from person_friend_map;")
	if tableCheck == nil {
		log.Println("Table person already exists")
		return nil
	}
	err := checkJson("data/PersonalData.json")
	if err != nil {
		return err
	}
	_, err = p.db.Exec(`CREATE TABLE person_friend_map(person_id varchar, friend_id varchar, CONSTRAINT person_id_FK FOREIGN KEY (person_id) REFERENCES people (id), CONSTRAINT friend_id_FK FOREIGN KEY (friend_id) REFERENCES friends (id))`)
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

func (p *Postgres) GetPeople() ([]model.Person, error) {
	var people []model.Person
	row, err := p.db.Query("SELECT * FROM people")
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer row.Close()

	for row.Next() {
		var a model.Person
		err = row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return nil, fmt.Errorf("row scan: %w", err)
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByID(id string) (model.Person, error) {
	row := p.db.QueryRow("SELECT * FROM people WHERE id = $1", id)
	if row.Err() != nil {
		return model.Person{}, row.Err()
	}
	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		return model.Person{}, err
	}
	return a, nil
}

func (p *Postgres) GetPersonByIndex(index int) (model.Person, error) {
	row := p.db.QueryRow("SELECT * FROM people WHERE index = $1", index)
	if row.Err() != nil {
		return model.Person{}, row.Err()
	}
	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		return model.Person{}, err
	}
	return a, nil
}

func (p *Postgres) GetPersonByGUID(guid string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE guid = $1", guid)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByIsActive(isActive bool) ([]model.Person, error) {
	var people []model.Person
	row, err := p.db.Query("SELECT * FROM people WHERE is_active = $1", isActive)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByBalance(balance string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE balance = $1", balance)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByAge(age int) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE age = $1", age)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByEyeColor(eyeColor string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE eye_color = $1", eyeColor)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByFirstName(firstname string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE name_first = $1", firstname)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByLastName(lastname string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE name_last = $1", lastname)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByGender(gender string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE gender = $1", gender)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByCompany(company string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE company = $1", company)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByEmail(email string) (model.Person, error) {
	row := p.db.QueryRow("SELECT * FROM people WHERE email = $1", email)
	if row.Err() != nil {
		return model.Person{}, row.Err()
	}

	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		return model.Person{}, err
	}
	return a, nil
}

func (p *Postgres) GetPersonByPhoneNumber(phoneNumber string) (model.Person, error) {
	row := p.db.QueryRow("SELECT * FROM people WHERE phone = $1", phoneNumber)
	if row.Err() != nil {
		return model.Person{}, row.Err()
	}

	var a model.Person

	err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
	if err != nil {
		return model.Person{}, err
	}
	return a, nil
}

func (p *Postgres) GetPersonByHouseNumber(houseNumber int) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE address_house_number = $1", houseNumber)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByStreetName(streetName string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE address_street = $1", streetName)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByZipCode(zipCode int) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE address_zip_code = $1", zipCode)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByCity(city string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE address_city = $1", city)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByState(state string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE address_state = $1", state)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByAbout(about string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE about = $1", about)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByRegistered(registered string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE registered = $1", registered)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByLatitude(latitude string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE latitude = $1", latitude)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}

func (p *Postgres) GetPersonByLongitude(longitude string) ([]model.Person, error) {
	var people []model.Person

	row, err := p.db.Query("SELECT * FROM people WHERE longitude = $1", longitude)
	if err != nil {
		return []model.Person{}, err
	}
	defer row.Close()
	for row.Next() {
		var a model.Person
		err := row.Scan(&a.ID, &a.Index, &a.GUID, &a.IsActive, &a.Balance, &a.Picture, &a.Age, &a.EyeColor, &a.Name.Firstname, &a.Name.Lastname, &a.Gender, &a.Company, &a.Email, &a.Phone, &a.Address.HouseNumber, &a.Address.Street, &a.Address.City, &a.Address.State, &a.Address.ZipCode, &a.About, &a.Registered, &a.Latitude, &a.Longitude)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, a)
	}
	return people, nil
}
