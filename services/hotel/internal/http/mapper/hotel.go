package mapper

import (
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/models"

	"github.com/google/uuid"
)

func HotelCreateRequestToEntity(req request.HotelCreate) models.HotelCreate {
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

func HotelUpdateRequestToEntity(req request.HotelUpdate) models.HotelUpdate {
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

func HotelEntityToResponse(req models.Hotel) response.Hotel {
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

func HotelShortEntityToShortResponse(req models.HotelShort) response.HotelShort {
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

func HotelUpdateEntityToResponse(id uuid.UUID, req models.HotelUpdate) response.HotelUpdate {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.HotelUpdate{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
	}
}
