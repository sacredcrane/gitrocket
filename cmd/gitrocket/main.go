package main

import (
	"fmt"

	"github.com/sacredcrane/gitrocket/internal/infra/config"
)

func main() {
	initServer()
}

func initServer() {
	cfg, err := config.LoadConfigFromFile()
	if err != nil {
		panic("Configuration missed!\n")
	}
	fmt.Printf("Config: %v", cfg)
}
