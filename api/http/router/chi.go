package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"auth_service/api/http/handler"
)

type Handlers struct {
	User handler.UserHandler
}

func NewHandlers(h *handler.Handler) *Handlers {
	return &Handlers{
		User: h,
	}
}

func New(r chi.Router, h *Handlers) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.User.Create)
		r.Get("/{id}", h.User.Get)
		r.Get("/", h.User.GetAll)
	})
}
