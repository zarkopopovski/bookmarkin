package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"crypto/sha1"
)

type DBConnection struct {
	db    *sql.DB
}

func OpenConnectionSession() (dbConnection *DBConnection) {
	dbConnection = new(DBConnection)
	dbConnection.createNewDBConnection()

	return
}

func (dbConnection *DBConnection) createNewDBConnection() (err error) {
	db, err := sql.Open("sqlite3", "./bookmarkin.db")
	if err != nil {
		panic(err)
	}

	fmt.Println("SQLite Connection is Active")
	dbConnection.db = db

	dbConnection.createDefaultUser()	

	return
}

func (dbConnection *DBConnection) createDefaultUser() bool {
	query := "SELECT id, username, password, email FROM users WHERE username='root' AND password='root'"

	err := dbConnection.db.QueryRow(query)

	if err != nil {
		sha1Hash := sha1.New()
		sha1Hash.Write([]byte("root"))
		sha1HashString := sha1Hash.Sum(nil)

		passwordEnc := fmt.Sprintf("%x", sha1HashString)
		
		query = "INSERT INTO users(id, username, password, email, date_created) VALUES('0','root','"+passwordEnc+"','0', date('now'))"
	
		_, err := dbConnection.db.Exec(query)

		if err != nil {
			return false
		}
	}

	return true
}