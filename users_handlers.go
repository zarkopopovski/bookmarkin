package main

import (
	"encoding/json"
	"net/http"
	//"fmt"
	//"strconv"
)

type UsersHandlers struct {
	dbConnection *DBConnection
}

func (uHandlers *UsersHandlers) CreateUserAccount(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	user := &User{
		Username:username,
		Password:password,
		Email:email}

	result := user.CreateNewUser(uHandlers.dbConnection)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(result); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}