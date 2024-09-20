package models

import (
	"database/sql"
	"log"

	"github.com/d11m08y03/digicup/config"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getDB() *sql.DB {
	if config.TestDB != nil {
		return config.TestDB
	}

	if config.DB == nil {
		log.Panic("Database connection is nil")
	}

	return config.DB
}

func CreateUser(user User) error {
	stmt, err := getDB().Prepare("INSERT INTO users(name, age, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Age, user.Email, user.Password)
	return err
}

func FindUserByEmail(email string) (*User, error) {
	row := getDB().QueryRow("SELECT id, name, age, email, password FROM users WHERE email = ?", email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	rows, err := getDB().Query("SELECT id, name, age, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
