package handler

import (
	"hotel/internal/repository/models"

	"github.com/google/uuid"

	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
)

type Service interface {
	HotelService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) HotelCreateRequestToEntity(req request.HotelCreate) models.HotelCreate {
	location := models.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return models.HotelCreate{
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
	}
}

func (h *Handler) HotelUpdateRequestToEntity(req request.HotelUpdate) models.HotelUpdate {
	location := models.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return models.HotelUpdate{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
	}
}

func (h *Handler) HotelEntityToResponse(req models.Hotel) response.Hotel {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.Hotel{
		ID:          req.ID,
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Address:     req.Address,
		Rating:      req.Rating,
		Location:    location,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
}

func (h *Handler) HotelShortEntityToShortResponse(req models.HotelShort) response.HotelShort {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.HotelShort{
		ID:       req.ID,
		Name:     req.Name,
		OwnerID:  req.OwnerID,
		Address:  req.Address,
		Rating:   req.Rating,
		Location: location,
	}
}

func (h *Handler) HotelUpdateEntityToResponse(id uuid.UUID, req models.HotelUpdate) response.HotelUpdate {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.HotelUpdate{
		ID:       id,
		Name:     req.Name,
		Address:  req.Address,
		Location: location,
	}
}
