package main

import (
	"finance/internal/api"
	"finance/pkg/configuration"
	"log"
)

func main() {
	config, err := configuration.NewConfig[api.Config]("configs", "appSettings", "Development")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	api.Run(config)
}
