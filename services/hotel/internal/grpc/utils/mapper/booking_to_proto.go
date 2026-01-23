package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
)

func locationResponseToProto(l *models.Location) *hotelv1.Location {
	return &hotelv1.Location{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

func HotelResponseToProto(resp *models.Hotel) *hotelv1.Hotel {
	return &hotelv1.Hotel{
		Id:          resp.ID.String(),
		Title:       resp.Title,
		Slug:        resp.Slug,
		OwnerId:     resp.OwnerID,
		Description: *resp.Description,
		Address:     resp.Address,
		Rating:      resp.Rating,
		Location:    locationResponseToProto(&resp.Location),
		CreatedAt:   timestamppb.New(resp.CreatedAt),
		UpdatedAt:   timestamppb.New(resp.UpdatedAt),
	}
}
