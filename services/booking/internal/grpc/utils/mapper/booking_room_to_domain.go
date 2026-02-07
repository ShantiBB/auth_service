package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/repository/models"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/utils/consts"
)

func CreateBookingRoomsToDomain(rooms []*bookingv1.CreateBookingRoomRequest) ([]*models.CreateBookingRoom, error) {
	result := make([]*models.CreateBookingRoom, len(rooms))

	for i, r := range rooms {
		roomID, err := uuid.Parse(r.RoomId)
		if err != nil {
			return nil, consts.ErrInvalidBookingRoomID
		}

		price, err := decimal.NewFromString(r.PricePerNight)
		if err != nil {
			return nil, consts.ErrInvalidPricePerNightID
		}

		result[i] = &models.CreateBookingRoom{
			RoomID:        roomID,
			Adults:        r.Adults,
			Children:      r.Children,
			PricePerNight: price,
		}
	}

	return result, nil
}
