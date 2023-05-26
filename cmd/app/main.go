package main

import (
	"go-template-wire/configs"
	"go-template-wire/internal/app"
	"log"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	app.Run(cfg)
}
