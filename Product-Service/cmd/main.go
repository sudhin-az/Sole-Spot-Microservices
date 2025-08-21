package main

import (
	"log"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}