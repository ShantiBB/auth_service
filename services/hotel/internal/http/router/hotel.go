package router

import (
	"github.com/go-chi/chi/v5"

	"hotel/internal/http/handler"
)

func hotelRouter(pattern string, r chi.Router, h *handler.Handler) {
	r.Route(pattern, func(r chi.Router) {
		r.Post("/", h.HotelCreate)
		r.Get("/", h.HotelGetAll)
		r.Get("/{id}", h.HotelGetByID)
		r.Put("/{id}", h.HotelUpdateByID)
		r.Delete("/{id}", h.HotelDeleteByID)
	})
}
