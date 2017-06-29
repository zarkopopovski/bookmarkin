package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
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
		Username: username,
		Password: passwordEnc,
		Email:    email}

	result := user.CreateNewUser(uHandlers.dbConnection)

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

func (uHandlers *UsersHandlers) LoginWithCredentials(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(password))
	sha1HashString := sha1Hash.Sum(nil)

	passwordEnc := fmt.Sprintf("%x", sha1HashString)

	user := &User{
		Username: username,
		Password: passwordEnc}

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

func (uHandlers *UsersHandlers) ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")
	password := r.FormValue("password")
	newPassword := r.FormValue("new_password")

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(password))
	sha1HashString := sha1Hash.Sum(nil)

	passwordEnc := fmt.Sprintf("%x", sha1HashString)

	user := &User{Id: userID}

	result := user.CheckUserByID(uHandlers.dbConnection)

	if result != nil {
		if result.Password == passwordEnc {
			sha1Hash := sha1.New()
			sha1Hash.Write([]byte(newPassword))
			sha1HashString := sha1Hash.Sum(nil)

			passwordEnc := fmt.Sprintf("%x", sha1HashString)

			fmt.Printf(newPassword + " " + passwordEnc)

			user := &User{Id: userID, Password: passwordEnc}

			testResult := user.UpdateUserPassword(uHandlers.dbConnection)

			if testResult {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)

				if err := json.NewEncoder(w).Encode(testResult); err != nil {
					panic(err)
				}
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
				panic(err)
			}

		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
				panic(err)
			}
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

}
