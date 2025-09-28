package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) sendError(w http.ResponseWriter, r *http.Request, code int, message interface{}) {
	w.WriteHeader(code)
	render.JSON(w, r, message)
}

func (h *Handler) sendJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	render.JSON(w, r, data)
}
