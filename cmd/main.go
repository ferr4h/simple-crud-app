package main

import (
	"log"
	"simple-crud-app/config"
	"simple-crud-app/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
