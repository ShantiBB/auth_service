package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/http/handler"
)

func roomRouter(pattern string, r chi.Router, h *handler.Handler) {
	r.Route(pattern, func(r chi.Router) {
		r.Post("/", h.RoomCreate)
		r.Get("/", h.RoomGetAll)
		r.Get("/{id}", h.RoomGetByID)
		r.Put("/{id}", h.RoomUpdateByID)
		r.Put("/{id}/update_status", h.RoomStatusUpdateByID)
		r.Delete("/{id}", h.RoomDeleteByID)
	})
}
