package auth

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"auth_service/api/http/handler"
	"auth_service/api/http/router"
	"auth_service/internal/config"
	"auth_service/internal/database/postgres"
	"auth_service/internal/service"
)

type App struct {
	Config *config.Config
}

func (app *App) MustLoad() {
	userRepo, err := postgres.NewRepository(app.Config)
	if err != nil {
		panic(err.Error())
	}

	userService := service.New(userRepo)
	userHandler := handler.New(userService)
	routerHandlers := router.NewHandlers(userHandler)

	r := chi.NewRouter()
	router.New(r, routerHandlers)

	server := fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port)
	fmt.Printf("Starting server on %s\n", server)
	if err = http.ListenAndServe(server, r); err != nil {
		panic(err.Error())
	}
}
