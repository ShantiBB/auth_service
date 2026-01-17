package mapper

import (
	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

func BookingRoomInfoToProto(r *models.BookingRoomFullInfo) *bookingv1.BookingRoomInfo {
	return &bookingv1.BookingRoomInfo{
		Id:            r.ID.String(),
		RoomId:        r.RoomID.String(),
		Adults:        uint32(r.Adults),
		Children:      uint32(r.Children),
		PricePerNight: r.PricePerNight.String(),
	}
}

func BookingRoomsInfoToProto(rooms []models.BookingRoomFullInfo) []*bookingv1.BookingRoomInfo {
	result := make([]*bookingv1.BookingRoomInfo, len(rooms))
	for i, r := range rooms {
		result[i] = BookingRoomInfoToProto(&r)
	}
	return result
}

func BookingRoomFullInfoToProto(r models.BookingRoomFullInfo) *bookingv1.BookingRoomFullInfo {
	return &bookingv1.BookingRoomFullInfo{
		Id:            r.ID.String(),
		RoomId:        r.RoomID.String(),
		Adults:        uint32(r.Adults),
		Children:      uint32(r.Children),
		PricePerNight: r.PricePerNight.String(),
		RoomLock:      RoomLockToProto(&r.RoomLock),
	}
}

func BookingRoomsFullInfoToProto(rooms []models.BookingRoomFullInfo) []*bookingv1.BookingRoomFullInfo {
	result := make([]*bookingv1.BookingRoomFullInfo, len(rooms))
	for i, r := range rooms {
		result[i] = BookingRoomFullInfoToProto(r)
	}
	return result
}
