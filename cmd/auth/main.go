package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"auth_service/internal/app/auth"
	"auth_service/internal/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env is not found: %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_FILENAME is not set")
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	authApp := auth.App{Config: cfg}
	authApp.MustLoad()
}
