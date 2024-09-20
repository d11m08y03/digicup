package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var TestDB *sql.DB

// Initializes an in-memory database for testing
func InitTestDB() {
	var err error
	TestDB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Failed to connect to the in-memory database: ", err)
	}

	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`

	_, err = TestDB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Failed to create users table in test: ", err)
	}

	log.Println("Test database successfully initialized!")
}

func CloseTestDB() {
	if err := TestDB.Close(); err != nil {
		log.Fatal("Failed to close the test database: ", err)
	}
}
