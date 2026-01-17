package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

func CreateBookingRoomsToDomain(rooms []*bookingv1.CreateBookingRoom) ([]models.CreateBookingRoom, error) {
	result := make([]models.CreateBookingRoom, len(rooms))

	for i, r := range rooms {
		roomID, err := uuid.Parse(r.RoomId)
		if err != nil {
			return nil, consts.ErrInvalidBookingRoomID
		}

		price, err := decimal.NewFromString(r.PricePerNight)
		if err != nil {
			return nil, consts.ErrInvalidPricePerNightID
		}

		result[i] = models.CreateBookingRoom{
			RoomID:        roomID,
			Adults:        uint8(r.Adults),
			Children:      uint8(r.Children),
			PricePerNight: price,
		}
	}

	return result, nil
}

func GetBookingRoomsRequestToDomain(req *bookingv1.GetBookingRoomsRequest) (uuid.UUID, error) {
	bookingID, err := uuid.Parse(req.BookingId)
	if err != nil {
		return uuid.UUID{}, consts.ErrInvalidBookingID
	}

	return bookingID, nil
}

func GetBookingRoomRequestToDomain(req *bookingv1.GetBookingRoomRequest) (uuid.UUID, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return uuid.UUID{}, consts.ErrInvalidBookingRoomID
	}

	return id, nil
}
