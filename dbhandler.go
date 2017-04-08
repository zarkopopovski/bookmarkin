package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
	db, err := sql.Open("sqlite3", "./selfmark.db")
	if err != nil {
		panic(err)
	}

	fmt.Println("SQLite Connection is Active")
	dbConnection.db = db

	return
}