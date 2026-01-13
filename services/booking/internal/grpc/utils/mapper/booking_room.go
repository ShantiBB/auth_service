package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

func CreateBookingRoomsToDomain(rooms []*bookingv1.CreateBookingRoom) ([]models.CreateBookingRoom, error) {
	result := make([]models.CreateBookingRoom, 0, len(rooms))

	for _, r := range rooms {
		roomID, err := uuid.Parse(r.RoomId)
		if err != nil {
			return nil, err
		}

		price, err := decimal.NewFromString(r.PricePerNight)
		if err != nil {
			return nil, err
		}

		result = append(
			result, models.CreateBookingRoom{
				RoomID:        roomID,
				Adults:        uint8(r.Adults),
				Children:      uint8(r.Children),
				PricePerNight: price,
			},
		)
	}

	return result, nil
}
