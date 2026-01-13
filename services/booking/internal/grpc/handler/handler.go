package handler

import (
	"buf.build/go/protovalidate"

	bookingv1 "booking/api/booking/v1"
)

type Service interface {
	BookingService
}

type Handler struct {
	bookingv1.UnimplementedBookingServiceServer
	svc       Service
	validator protovalidate.Validator
}

func New(svc Service, validator protovalidate.Validator) *Handler {
	return &Handler{svc: svc, validator: validator}
}
