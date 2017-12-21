package main

import (
	"fmt"
	"net/http"

	"io/ioutil"
)

type ApiConnection struct {
	dbConnection *DBConnection
	bHandlers    *BookmarkHandlers
	uHandlers    *UsersHandlers
}

func CreateApiConnection(config *Config) *ApiConnection {
	API := &ApiConnection{
		dbConnection: OpenConnectionSession(),
		bHandlers:    &BookmarkHandlers{},
		uHandlers:    &UsersHandlers{},
	}
	API.bHandlers.dbConnection = API.dbConnection
	API.uHandlers.dbConnection = API.dbConnection

	return API
}

func (c *ApiConnection) Index(w http.ResponseWriter, r *http.Request) {
	index, err := ioutil.ReadFile("./web/index.html")

	panic(err)

	if err != nil {
		panic(err)
		return
	}

	fmt.Fprintf(w, string(index))
}
