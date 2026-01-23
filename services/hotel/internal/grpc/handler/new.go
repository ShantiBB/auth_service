package handler

import (
	"context"

	"buf.build/go/protovalidate"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
)

type HotelService interface {
	CreateHotel(ctx context.Context, h *models.CreateHotel) (*models.Hotel, error)
	HotelGetBySlug(ctx context.Context, hotel models.HotelRef) (models.Hotel, error)
	GetHotels(ctx context.Context, hotelInfo *models.HotelList, sortField string, page, limit uint64) (
		*models.HotelList, error,
	)
	HotelUpdateBySlug(ctx context.Context, hotel models.HotelRef, h models.UpdateHotel) error
	HotelTitleUpdateBySlug(
		ctx context.Context, hotel models.HotelRef, h models.UpdateHotelTitle,
	) (models.UpdateHotelTitle, error)
	HotelDeleteBySlug(ctx context.Context, hotel models.HotelRef) error
}

type Service interface {
	HotelService
}

type Handler struct {
	hotelv1.UnimplementedHotelServiceServer
	svc       Service
	validator protovalidate.Validator
}

func New(svc Service, validator protovalidate.Validator) *Handler {
	return &Handler{svc: svc, validator: validator}
}
