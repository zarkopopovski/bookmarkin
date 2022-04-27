package main

import (
	"fmt"
	"net/http"

	"io/ioutil"
	
	"github.com/julienschmidt/httprouter"
)

type ApiConnection struct {
	dbConnection *DBConnection
	bHandlers    *BookmarkHandlers
	uHandlers    *UsersHandlers
}

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
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

func (c *ApiConnection) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	index, err := ioutil.ReadFile("./web/index.html")

	//panic(err)

	if err != nil {
			panic(err)
			return
	}

	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, string(index))
}
