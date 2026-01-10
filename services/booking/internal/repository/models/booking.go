package models

import "time"

type BookingCreate struct {
	UserID      int64
	HotelID     string
	CheckIn     string
	CheckOut    string
	GuestName   string
	GuestEmail  *string
	GuestPhone  *string
	Currency    string
	TotalAmount string
}

type BookingUpdate struct {
	CheckIn     *string
	CheckOut    *string
	GuestName   *string
	GuestEmail  *string
	GuestPhone  *string
	TotalAmount *string
}

type BookingStatusUpdate struct {
	Status BookingStatus
}

type Booking struct {
	ID          string
	UserID      int64
	HotelID     string
	CheckIn     string
	CheckOut    string
	Status      BookingStatus
	GuestName   string
	GuestEmail  *string
	GuestPhone  *string
	Currency    string
	TotalAmount string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
