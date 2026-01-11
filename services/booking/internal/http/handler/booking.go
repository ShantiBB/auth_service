package handler

import (
	"context"
	"net/http"

	"booking/internal/http/dto/request"
	_ "booking/internal/http/dto/response"
	"booking/internal/http/utils/helper"
	"booking/internal/http/utils/mapper"
	"booking/internal/http/utils/validation"
	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

type BookingService interface {
	BookingCreate(ctx context.Context, b models.CreateBooking, rooms []models.CreateBookingRoom) (models.Booking, error)
}

// BookingCreate godoc
// @Summary      Create booking
// @Description  Create a new booking with rooms
// @Tags         bookings
// @Accept       json
// @Produce      json
// @Param        request  body      request.CreateBooking  true  "Booking data with rooms"
// @Success      201      {object}  response.Booking
// @Failure      400      {object}  response.ErrorSchema
// @Failure      401      {object}  response.ErrorSchema
// @Failure      409      {object}  response.ErrorSchema
// @Failure      500      {object}  response.ErrorSchema
// @Security     Bearer
// @Router       /bookings [post]
func (h *Handler) BookingCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.CreateBooking
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	newBooking := mapper.BookingCreateRequestToEntity(req)
	rooms := mapper.BookingRoomsCreateRequestToEntity(req.Rooms)

	createdBooking, err := h.svc.BookingCreate(ctx, newBooking, rooms)
	errHandler := &helper.ErrorHandler{
		Conflict:   consts.RoomLockAlreadyExist,
		BadRequest: consts.ErrPriceChanged,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	bookingResponse := mapper.BookingCreateEntityToResponse(createdBooking)
	helper.SendSuccess(w, r, http.StatusCreated, bookingResponse)
}
