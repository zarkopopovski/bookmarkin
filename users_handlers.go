package main

import (
	"encoding/json"
	"net/http"
	"crypto/sha1"
	"fmt"
	//"strconv"
)

type UsersHandlers struct {
	dbConnection *DBConnection
}

func (uHandlers *UsersHandlers) CreateUserAccount(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(password))
	sha1HashString := sha1Hash.Sum(nil)

	passwordEnc := fmt.Sprintf("%x", sha1HashString)

	user := &User{
		Username:username,
		Password:passwordEnc,
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

func (uHandlers *UsersHandlers) LoginWithCredentials(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(password))
	sha1HashString := sha1Hash.Sum(nil)

	passwordEnc := fmt.Sprintf("%x", sha1HashString)

	user := &User{
		Username:username,
		Password:passwordEnc}

	result := user.CheckUserCredentials(uHandlers.dbConnection)

	if result != nil {
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