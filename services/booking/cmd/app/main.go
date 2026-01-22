package main

import (
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"booking/internal/app/booking"
	"booking/internal/config"
	"booking/pkg/lib/logger"
)

//	@title			Swagger Bookings API
//	@version		1.0
//	@description	Bookings service for microservices.

//	@host		localhost:8083
//	@BasePath	/api/v1

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
func main() {
	if err := cleanenv.ReadConfig(".env", &struct{}{}); err != nil {
		slog.Warn("failed to load env", "error", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is not set")
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	log := logger.New(cfg.Env, cfg.LogLevel)
	bookingApp := booking.App{
		Config: cfg,
		Logger: log,
	}
	bookingApp.MustLoadGRPC()
}
