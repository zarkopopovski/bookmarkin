package main

import (
	"log"
	"net/http"
)

func main() {

	apiConnection := CreateApiConnection()

	router := NewRouter(apiConnection)

	log.Fatal(http.ListenAndServe(":8080", router))

}
