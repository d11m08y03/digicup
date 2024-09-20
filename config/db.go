package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./digicup.db")
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`
	
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Failed to create users table: ", err)
	}

	log.Println("Database successfully initialized!")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("Failed to close the database: ", err)
	}
}
