package mapper

import (
	"github.com/shopspring/decimal"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
)

func CreateRoomRequestToDomain(req *hotelv1.CreateRoomRequest) *models.CreateRoom {
	room := &models.CreateRoom{
		Description: req.Description,
		Title:       req.Title,
		RoomNumber:  req.RoomNumber,
		Type:        roomTypeToDomain(req.Type),
		Price:       decimal.NewFromFloat(float64(req.Price)),
		Amenities:   req.Amenities,
		Images:      req.Images,
		Capacity:    int(req.Capacity),
		AreaSqm:     float64(req.AreaSqm),
		Floor:       int(req.Floor),
	}

	return room
}

func roomTypeToDomain(status hotelv1.RoomType) models.RoomType {
	var s models.RoomType
	switch status {
	case hotelv1.RoomType_ROOM_TYPE_SINGLE:
		s = models.RoomTypeSingle
	case hotelv1.RoomType_ROOM_TYPE_DOUBLE:
		s = models.RoomTypeDouble
	case hotelv1.RoomType_ROOM_TYPE_SUITE:
		s = models.RoomTypeSuite
	case hotelv1.RoomType_ROOM_TYPE_FAMILY:
		s = models.RoomTypeFamily
	case hotelv1.RoomType_ROOM_TYPE_PRESIDENTIAL:
		s = models.RoomTypePresidential
	default:
		s = models.RoomTypeUnspecified
	}
	return s
}
