package mapper

import (
	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

func BookingRoomToProto(r *models.BookingRoomInfo) *bookingv1.BookingRoomResponse {
	return &bookingv1.BookingRoomResponse{
		RoomId:        r.RoomID.String(),
		Adults:        uint32(r.Adults),
		Children:      uint32(r.Children),
		PricePerNight: r.PricePerNight.String(),
	}
}

func BookingRoomsInfoToProto(rooms []models.BookingRoomInfo) []*bookingv1.BookingRoomInfo {
	result := make([]*bookingv1.BookingRoomInfo, 0, len(rooms))
	for _, r := range rooms {
		result = append(
			result, &bookingv1.BookingRoomInfo{
				Id:            r.ID.String(),
				RoomId:        r.RoomID.String(),
				Adults:        uint32(r.Adults),
				Children:      uint32(r.Children),
				PricePerNight: r.PricePerNight.String(),
			},
		)
	}
	return result
}
