package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() (con *sql.DB, err error) {

	db, err := sql.Open("sqlite3", "file:youtube.db?cache=share&mode=rwc&_journal_mode=WAL&synchronous=normal&temp_store=memory&mmap_size=30000000000&page_size=32768")

	if err != nil {
		return nil, err
	}
	return db, nil

}
