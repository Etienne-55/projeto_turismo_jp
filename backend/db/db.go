package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


var DB *sql.DB

func InitDB() {
	var err error 
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()
}

func createTables() {

	createTouristTable := `
	CREATE TABLE IF NOT EXISTS tourist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createTouristTable)

	if err != nil {
		panic("could not create tourist table")
	}
}

