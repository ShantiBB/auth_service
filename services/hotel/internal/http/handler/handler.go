package handler

import (
	"github.com/go-playground/validator/v10"

	"fukuro-reserve/pkg/utils/consts"
)

type Service interface {
	HotelService
	RoomService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) customValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return consts.FieldRequired.Error()
	default:
		return err.Error()
	}
}
