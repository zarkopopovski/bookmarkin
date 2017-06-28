package main

import (
	"fmt"
	"net/http"

	"io/ioutil"

	"git.cerebralab.com/george/logo"
)

type ApiConnection struct {
	dbConnection *DBConnection
	bHandlers    *BookmarkHandlers
	uHandlers    *UsersHandlers
}

func CreateApiConnection() *ApiConnection {
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

	logo.RuntimeError(err)

	if err != nil {
		logo.RuntimeError(err)
		return
	}

	fmt.Fprintf(w, string(index))
}
