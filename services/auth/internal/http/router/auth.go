package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/http/handler"
)

func authRouter(pattern string, r chi.Router, h *handler.Handler) {
	r.Route(pattern, func(r chi.Router) {
		r.Post("/register", h.RegisterByEmail)
		r.Post("/refresh", h.RefreshToken)
		r.Post("/login", h.LoginByEmail)
	})
}
