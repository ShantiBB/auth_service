package mapper

import (
	"github.com/google/uuid"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
	"hotel/pkg/lib/utils/consts"
)

func locationRequestToDomain(req *hotelv1.CreateBookingLocationRequest) models.Location {
	return models.Location{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}

func CreateHotelRequestToDomain(req *hotelv1.CreateHotelRequest) *models.CreateHotel {
	return &models.CreateHotel{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		Title:       req.Title,
		OwnerID:     req.OwnerId,
		Description: req.Description,
		Address:     req.Address,
		Location:    locationRequestToDomain(req.Location),
	}
}

func GetHotelsRequestToDomain(req *hotelv1.GetHotelsRequest) (uint64, uint64, models.HotelRef) {
	hotelInfo := models.HotelRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
	}
	return req.Page, req.Limit, hotelInfo
}

func GetHotelRequestToDomain(idStr string) (uuid.UUID, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.UUID{}, consts.ErrInvalidHotelID
	}

	return id, nil
}
