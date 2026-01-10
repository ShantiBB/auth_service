package booking

import (
	"fmt"
	"log/slog"
	"net/http"

	"booking/internal/config"
	"booking/internal/http/handler"
	"booking/internal/http/router"
	"booking/internal/repository/postgres"
	"booking/internal/service"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Config *config.Config
}

func (app *App) MustLoad() {
	repo, err := postgres.NewRepository(app.Config)
	if err != nil {
		panic(err.Error())
	}

	svc := service.New(repo)
	h := handler.New(svc)

	r := chi.NewRouter()
	router.New(r, h)

	server := fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port)
	slog.Info("Starting server", "address", server)
	if err = http.ListenAndServe(server, r); err != nil {
		panic(err.Error())
	}
}
