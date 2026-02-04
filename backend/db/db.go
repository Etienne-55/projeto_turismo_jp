package db

import (
	"database/sql"
	"log"
	"projeto_turismo_jp/utils"

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
	createAdminUser()
}

func createTables() {
	createTouristTable := `
	CREATE TABLE IF NOT EXISTS tourist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT DEFAULT 'user'
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
	trip_review TEXT DEFAULT '',
	status TEXT DEFAULT 'upcoming' CHECK(status IN ('upcoming', 'ongoing', 'completed')),
	tourist_id INTEGER NOT NULL,
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

func createAdminUser() {
    var count int
    err := DB.QueryRow("SELECT COUNT(*) FROM tourist WHERE role = 'admin'").Scan(&count)
    if err != nil || count > 0 {
        return 
    }
    
    hashedPassword, _ := utils.HashPassword("admin123")  
    
    query := `INSERT INTO tourist (email, password, role) VALUES (?, ?, ?)`
    _, err = DB.Exec(query, "admin@proton.me", hashedPassword, "admin")
    
    if err != nil {
        log.Printf("Failed to create admin: %v", err)
    } else {
        log.Println("Admin user created: admin@example.com")
    }
}
