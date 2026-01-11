package router

import (
	"github.com/go-chi/chi/v5"

	"booking/internal/http/handler"
)

func bookingRouter(pattern string, r chi.Router, h *handler.Handler) {
	r.Route(
		pattern, func(r chi.Router) {
			r.Post("/", h.BookingCreate)
		},
	)
}
