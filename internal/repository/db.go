package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type Postgres struct {
	Db *sql.DB
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

	return &Postgres{Db: openedDB}, err
}
