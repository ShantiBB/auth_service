package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Latitude  float32
	Longitude float32
}

type CreateHotel struct {
	CountryCode string
	CitySlug    string
	Title       string
	Slug        string
	OwnerID     int64
	Description *string
	Address     string
	Location    Location
}

type UpdateHotel struct {
	Description *string
	Address     string
	Location    Location
}

type UpdateHotelTitle struct {
	Title string
	Slug  string
}

type HotelShort struct {
	ID       uuid.UUID
	Title    string
	Slug     string
	OwnerID  int64
	Address  string
	Rating   *float32
	Location Location
}

type Hotel struct {
	ID          uuid.UUID
	Title       string
	Slug        string
	OwnerID     int64
	Description *string
	Address     string
	Rating      *float32
	Location    Location
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type HotelList struct {
	CountryCode string
	CitySlug    string
	Hotels      []*HotelShort
	TotalCount  uint64
}

func (h *CreateHotel) ToRead() Hotel {
	return Hotel{
		Title:       h.Title,
		Slug:        h.Slug,
		OwnerID:     h.OwnerID,
		Description: h.Description,
		Address:     h.Address,
		Location:    h.Location,
	}
}

type HotelRef struct {
	CountryCode string
	CitySlug    string
	HotelSlug   string
}
