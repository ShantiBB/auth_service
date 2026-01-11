package mapper

import (
	"booking/internal/http/dto/request"
	"booking/internal/repository/models"
)

func BookingRoomsCreateRequestToEntity(rooms []request.CreateBookingRoom) []models.CreateBookingRoom {
	var roomsEntity []models.CreateBookingRoom
	for _, room := range rooms {
		roomsEntity = append(
			roomsEntity,
			models.CreateBookingRoom{
				RoomID:        room.RoomID,
				Adults:        room.Adults,
				Children:      room.Children,
				PricePerNight: room.PricePerNight,
			},
		)
	}

	return roomsEntity
}
