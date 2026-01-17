package consts

import "errors"

// Bookings
var (
	ErrBookingNotFound      = errors.New("booking not found")
	ErrBookingRoomNotFound  = errors.New("booking room not found")
	ErrRoomLockNotFound     = errors.New("room lock room not found")
	ErrRoomLockAlreadyExist = errors.New("room lock already exists")
	ErrInvalidDates         = errors.New("invalid booking dates")
	ErrPriceChanged         = errors.New("expected total amount does not match calculated total")
)

var (
	ErrInvalidHotelID               = errors.New("invalid hotel ID")
	ErrInvalidBookingID             = errors.New("invalid booking ID")
	ErrInvalidBookingRoomID         = errors.New("invalid booking room ID")
	ErrInvalidPricePerNightID       = errors.New("invalid price per night. example: 123.45")
	ErrInvalidExpectedTotalAmountID = errors.New("invalid expected total amount. example: 123.45")
	ErrInternalServer               = errors.New("internal server error")
	ErrInvalidRequest               = errors.New("invalid request body")
)
