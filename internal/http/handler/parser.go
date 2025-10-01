package handler

import (
	"net/http"

	"github.com/go-chi/render"

	"auth_service/internal/http/lib/schemas"
	"auth_service/internal/http/lib/validation"
)

func (h *Handler) ParseJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		errMsg := schemas.NewErrorResponse("Invalid JSON body")
		h.sendError(w, r, http.StatusBadRequest, errMsg)
		return false
	}

	if errResp := validation.CheckErrors(v); errResp != nil {
		h.sendError(w, r, http.StatusBadRequest, errResp)
		return false
	}

	return true
}
