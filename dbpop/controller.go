package dbpop

import (
	"Gin-Postgres-API/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "TopSecret123!"
	dbname   = "postgres"
)

func Populate() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// JSON-Datei lesen
	file, err := os.ReadFile("./data/PersonalData.json")
	if err != nil {
		log.Fatal(err)
	}

	// Struktur für Daten
	var data []utils.Person
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Daten in die Datenbank einfügen
	for _, person := range data {
		insertPerson(db, person)
		insertFriends(db, person)
		insertMap(db, person)
	}

}

func insertPerson(db *sql.DB, p utils.Person) {
	insertStatement := `
		INSERT INTO person (
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

func insertFriends(db *sql.DB, p utils.Person) {

	insertStatement := `INSERT INTO friend (id, firstname, lastname) VALUES ($1, $2, $3)`

	for i := 0; i < len(p.Friends); i++ {
		_, err := db.Exec(insertStatement, p.Friends[i].ID, p.Friends[i].Name.Firstname, p.Friends[i].Name.Lastname)
		if err != nil {
			log.Printf("Error inserting person: %v\n", err)
		}
	}
}

func insertMap(db *sql.DB, p utils.Person) {

	insertStatement := `INSERT INTO mapping_friend_person (personid, friendid) VALUES ($1, $2)`

	for i := 0; i < len(p.Friends); i++ {
		_, err := db.Exec(insertStatement, p.ID, p.Friends[i].ID)
		if err != nil {
			log.Printf("Error inserting person: %v\n", err)
		}
	}
}
