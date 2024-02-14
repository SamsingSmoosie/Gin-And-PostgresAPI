package dbpop

import (
	"Gin-Postgres-API/internal/model"
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

// Create Creates and populates the db if nonexistent
func Create(filepath string, db *sql.DB) error {
	var people []model.Person

	//-----To be changed to reference the db snapshot not the data in memory------
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &people)
	if err != nil {
		return err
	}
	//----------------------------------------------------------------------------

	CreatePeople(db, people)
	CreateFriends(db, people)
	CreateMap(db, people)

	return nil
}

// CreatePeople Check if table "person" exists. If not it will create the table and populate it
func CreatePeople(db *sql.DB, people []model.Person) error {
	_, tableCheck := db.Query("select * from people;")
	if tableCheck == nil {
		log.Println("Table people already exists")
		return nil
	}

	_, err := db.Exec("CREATE TABLE people(id varchar primary key, index int, guid varchar, is_active bool, balance varchar, picture varchar, age int, eye_color varchar, name_first varchar, name_last varchar, gender varchar, company varchar, email varchar, phone varchar, address_house_number int, address_street varchar, address_city varchar, address_state varchar, address_zip_code int, about varchar, registered varchar, latitude float8, longitude float8)")
	if err != nil {
		return err
	}
	for _, person := range people {
		insertPerson(db, person)
	}
	log.Println("Table people created and populated")
	return nil
}

// CreateFriends Check if table "friend" exists. If not it will create the table and populate it
func CreateFriends(db *sql.DB, people []model.Person) error {
	_, tableCheck := db.Query("select * from friends;")
	if tableCheck == nil {
		log.Println("Table friends already exists")
		return nil
	}
	_, err := db.Exec("CREATE TABLE friends(id varchar primary key, name_first varchar, name_last varchar)")
	if err != nil {
		return err
	}
	for _, person := range people {
		insertFriends(db, person)
	}
	log.Println("Table friend created and populated")
	return nil
}

// CreateMap Check if table "person_friend_map" exists. If not it will create the table and populate it
func CreateMap(db *sql.DB, people []model.Person) error {
	_, tableCheck := db.Query("select * from person_friend_map;")
	if tableCheck == nil {
		log.Println("Table person already exists")
		return nil
	}
	_, err := db.Exec(`CREATE TABLE person_friend_map(person_id varchar, friend_id varchar, CONSTRAINT person_id_FK FOREIGN KEY (person_id) REFERENCES people (id), CONSTRAINT friend_id_FK FOREIGN KEY (friend_id) REFERENCES friends (id))`)
	if err != nil {
		return err
	}
	for _, person := range people {
		insertMap(db, person)
	}
	log.Println("Table person_friend_map created and populated")
	return nil
}

func insertPerson(db *sql.DB, p model.Person) {
	insertStatement := `
		INSERT INTO people (
			id, index, guid, is_active, balance, picture, age, eye_color,
			name_first, name_last, gender, company, email, phone,
			address_house_number, address_street, address_city,
			address_state, address_zip_code, about, registered,
			latitude, longitude
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
	`
	_, err := db.Exec(insertStatement,
		p.ID, p.Index, p.GUID, p.IsActive, p.Balance, p.Picture, p.Age, p.EyeColor,
		p.Name.Firstname, p.Name.Lastname, p.Gender, p.Company, p.Email, p.Phone,
		p.Address.HouseNumber, p.Address.Street, p.Address.City,
		p.Address.State, p.Address.ZipCode, p.About, p.Registered,
		p.Latitude, p.Longitude,
	)

	if err != nil {
		log.Printf("Error inserting person: %v\n", err)
	}
}

func insertFriends(db *sql.DB, p model.Person) {

	insertStatement := `INSERT INTO friends (id, name_first, name_last) VALUES ($1, $2, $3)`

	for _, friend := range p.Friends {
		_, err := db.Exec(insertStatement, friend.ID, friend.Name.Firstname, friend.Name.Lastname)
		if err != nil {
			log.Printf("Error inserting friends: %v\n", err)
		}
	}
}

func insertMap(db *sql.DB, p model.Person) {

	insertStatement := `INSERT INTO person_friend_map (person_id, friend_id) VALUES ($1, $2)`

	for _, friend := range p.Friends {
		_, err := db.Exec(insertStatement, p.ID, friend.ID)
		if err != nil {
			log.Printf("Error inserting person: %v\n", err)
		}
	}
}
