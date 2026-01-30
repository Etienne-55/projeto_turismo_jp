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

	createTripTable := `
	CREATE TABLE IF NOT EXISTS trip (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	lodging_location TEXT NOT NULL,
	trip_description TEXT NOT NULL,
	arrival_date DATE NOT NULL,
	departure_date DATE NOT NULL,
	status TEXT DEFAULT 'upcoming' CHECK(status IN ('upcoming', 'ongoing', 'completed')),
	tourist_id INTEGER,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(tourist_id) REFERENCES tourist(id)
	)
	`
	_, err = DB.Exec(createTripTable)

	if err != nil {
		panic("could not create trip table")
	}
}

