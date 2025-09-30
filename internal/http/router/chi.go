package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"auth_service/internal/http/handler"
)

func New(r chi.Router, h *handler.Handler) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.UserCreate)
		r.Get("/", h.UserList)
		r.Get("/search", h.UserGetByUsernameOrEmail)
		r.Get("/{id}", h.UserGetByID)
		r.Put("/{id}", h.UserUpdateByID)
		r.Delete("/{id}", h.UserDeleteByID)
	})
}
