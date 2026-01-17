package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateBookingRoom struct {
	PricePerNight decimal.Decimal
	BookingID     uuid.UUID
	RoomID        uuid.UUID
	Adults        uint8
	Children      uint8
}

type BookingRoomGuestCounts struct {
	Adults   uint8
	Children uint8
}

type BookingRoomFullInfo struct {
	CreatedAt     time.Time
	PricePerNight decimal.Decimal
	RoomLock      RoomLockShort
	ID            uuid.UUID
	BookingID     uuid.UUID
	RoomID        uuid.UUID
	Adults        uint8
	Children      uint8
}

func (b *CreateBookingRoom) ToRead() BookingRoomFullInfo {
	return BookingRoomFullInfo{
		BookingID:     b.BookingID,
		RoomID:        b.RoomID,
		Adults:        b.Adults,
		Children:      b.Children,
		PricePerNight: b.PricePerNight,
	}
}
