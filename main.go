package main

import (
	"log"
	"net/http"
)

//TODO: Implement configuration by using config file
func main() {

	apiConnection := CreateApiConnection()

	router := NewRouter(apiConnection)

	log.Fatal(http.ListenAndServe(":8080", router))

}
