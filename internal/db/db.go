package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 8181
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
)

func Open() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}
