package models

import (
	"time"

	"github.com/google/uuid"
)

type HotelCreate struct {
	Name        string
	OwnerID     int64
	Description *string
	Address     string
	Location    Location
}

type Hotel struct {
	ID          uuid.UUID
	Name        string
	OwnerID     int64
	Description *string
	Address     string
	Location    Location
	Rating      *float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func (h *HotelCreate) ToRead() Hotel {
	return Hotel{
		Name:        h.Name,
		OwnerID:     h.OwnerID,
		Description: h.Description,
		Address:     h.Address,
		Location:    h.Location,
	}
}
