package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/knq/ini"
)

type Config struct {
	mailServer   string
	mailPort     int
	mailUsername string
	mailPassword string
}

func main() {
	fileCfg, err := ini.LoadFile("config.cfg")
	if err != nil {
		log.Fatal("Error with service configuration %s", err)
	}

	port := fileCfg.GetKey("service.port")

	if port == "" {
		log.Fatal("Error with port number configuration")
	}

	serverPort := fileCfg.GetKey("service.mailport")
	serverPortI, _ := strconv.Atoi(serverPort)

	config := &Config{
		mailServer:   fileCfg.GetKey("service.mailserver"),
		mailPort:     serverPortI,
		mailUsername: fileCfg.GetKey("service.mailusername"),
		mailPassword: fileCfg.GetKey("service.mailpassword"),
	}

	apiConnection := CreateApiConnection(config)

	router := NewRouter(apiConnection)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
