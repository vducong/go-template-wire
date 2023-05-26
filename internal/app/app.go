package app

import (
	"go-template-wire/configs"
	"log"
)

func Run(cfg *configs.Config) {
	s, err := initDeps(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies of server: %v\n", err)
		return
	}

	s.Start()
}
