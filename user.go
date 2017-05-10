package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) CreateNewUser(dbConnection *DBConnection) bool {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(time.Now().String() + user.Username + user.Password + user.Email))
	sha1HashString := sha1Hash.Sum(nil)

	userID := fmt.Sprintf("%x", sha1HashString)

	query := "INSERT INTO users(id, username, password, email, date_created) VALUES('" + userID + "','" + user.Username + "','" + user.Password + "','" + user.Email + "', date('now'))"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (user *User) CheckUserCredentials(dbConnection *DBConnection) *User {
	query := "SELECT id, username, password, email FROM users WHERE username='" + user.Username + "' AND password='" + user.Password + "'"

	newUser := new(User)

	err := dbConnection.db.QueryRow(query).Scan(
		&newUser.Id,
		&newUser.Username,
		&newUser.Password,
		&newUser.Email)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return newUser
}

func (user *User) CheckUserByID(dbConnection *DBConnection) *User {
	query := "SELECT id, username, password, email FROM users WHERE id='" + user.Id + "'"

	newUser := new(User)

	err := dbConnection.db.QueryRow(query).Scan(
		&newUser.Id,
		&newUser.Username,
		&newUser.Password,
		&newUser.Email)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return newUser
}

func (user *User) UpdateUserPassword(dbConnection *DBConnection) bool {
	query := "UPDATE users SET password='" + user.Password + "' WHERE id='" + user.Id + "'"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
