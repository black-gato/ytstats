package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() (con *sql.DB, err error) {

	db, err := sql.Open("sqlite3", "./youtube.db")

	if err != nil {
		return nil, err
	}
	return db, nil

}
