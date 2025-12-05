package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"hotel/internal/app/hotel"
	"hotel/internal/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("failed load env", "error", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_FILENAME is not set")
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	hotelApp := hotel.App{Config: cfg}
	hotelApp.MustLoad()
}
