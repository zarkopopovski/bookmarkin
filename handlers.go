package main

import (
	"fmt"
	"net/http"
)

type ApiConnection struct {
	dbConnection *DBConnection
	bHandlers    *BookmarkHandlers
}

func CreateApiConnection() *ApiConnection {
	API := &ApiConnection{
		dbConnection: OpenConnectionSession(),
		bHandlers:    &BookmarkHandlers{},
	}
	API.bHandlers.dbConnection = API.dbConnection

	return API
}

func (c *ApiConnection) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
