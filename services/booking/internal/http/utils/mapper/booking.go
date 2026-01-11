package mapper

import (
	"booking/internal/http/dto/request"
	"booking/internal/http/dto/response"
	"booking/internal/repository/models"
)

func BookingCreateRequestToEntity(req request.CreateBooking) models.CreateBooking {
	return models.CreateBooking{
		UserID:              req.UserID,
		HotelID:             req.HotelID,
		CheckIn:             req.CheckIn,
		CheckOut:            req.CheckOut,
		GuestName:           req.GuestName,
		GuestEmail:          req.GuestEmail,
		GuestPhone:          req.GuestPhone,
		Currency:            req.Currency,
		ExpectedTotalAmount: req.ExpectedTotalAmount,
	}
}

func BookingCreateEntityToResponse(b models.Booking) response.Booking {
	return response.Booking{
		ID:               b.ID,
		UserID:           b.UserID,
		HotelID:          b.HotelID,
		CheckIn:          b.CheckIn,
		CheckOut:         b.CheckOut,
		Status:           b.Status,
		GuestName:        b.GuestName,
		GuestEmail:       b.GuestEmail,
		GuestPhone:       b.GuestPhone,
		Currency:         b.Currency,
		FinalTotalAmount: b.FinalTotalAmount,
		CreatedAt:        b.CreatedAt,
		UpdatedAt:        b.UpdatedAt,
	}
}
