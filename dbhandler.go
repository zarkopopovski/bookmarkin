package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DBConnection struct {
	db *sql.DB
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

	dbConnection.setupInitialDatabase()
	dbConnection.createDefaultData()

	return
}

func (dbConnection *DBConnection) setupInitialDatabase() (err error) {
	statement, _ := dbConnection.db.Prepare("CREATE TABLE IF NOT EXISTS users (id VARCHAR PRIMARY KEY, username VARCHAR, email VARCHAR, password VARCHAR, date_created VARCHAR)")
	statement.Exec()

	statement, _ = dbConnection.db.Prepare("CREATE TABLE IF NOT EXISTS groups (id VARCHAR PRIMARY KEY, user_id VARCHAR, group_name VARCHAR)")
	statement.Exec()

	statement, _ = dbConnection.db.Prepare("CREATE TABLE IF NOT EXISTS bookmarks (id VARCHAR PRIMARY KEY, user_id VARCHAR, bookmark_url VARCHAR, bookmark_title VARCHAR, bookmark_icon VARCHAR, bookmark_group VARCHAR)")
	statement.Exec()

	return
}

func (dbConnection *DBConnection) createDefaultData() bool {
	query := "SELECT id, username, password, email FROM users WHERE username='root' AND password='root'"

	err := dbConnection.db.QueryRow(query)

	if err != nil {
		sha1Hash := sha1.New()
		sha1Hash.Write([]byte("root"))
		sha1HashString := sha1Hash.Sum(nil)

		passwordEnc := fmt.Sprintf("%x", sha1HashString)

		query = "INSERT INTO users(id, username, password, email, date_created) VALUES('11','root','" + passwordEnc + "','0', date('now'))"

		_, err := dbConnection.db.Exec(query)

		if err != nil {
			return false
		}

		query = "INSERT INTO groups(id, user_id, group_name) VALUES('111','11','Default')"

		_, err = dbConnection.db.Exec(query)

		if err != nil {
			return false
		}
	}

	return true
}
