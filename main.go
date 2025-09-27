package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"auth_service/api/http/handler"
	"auth_service/api/http/router"
	"auth_service/internal/database/postgres"
	"auth_service/internal/service"
)

func main() {
	sdn := "user=sentinel password=1221 dbname=auth sslmode=disable"

	userRepo, err := postgres.NewRepository(sdn)
	if err != nil {
		panic(err.Error())
	}

	userService := service.New(userRepo)
	userHandler := handler.New(userService)

	r := chi.NewRouter()
	router.New(r, userHandler)

	fmt.Println("Starting server on :8080")
	if err = http.ListenAndServe(":8080", r); err != nil {
		panic(err.Error())
	}

	panic("unreachable")
}
