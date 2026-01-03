package response

import (
	"hotel/internal/http/dto"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Room struct {
	ID          uuid.UUID       `json:"id"`
	HotelID     uuid.UUID       `json:"hotel_id"`
	Description string          `json:"description"`
	RoomNumber  string          `json:"room_number"`
	Type        dto.RoomType    `json:"type"`
	Status      dto.RoomStatus  `json:"status"`
	Price       decimal.Decimal `json:"price"`
	Capacity    int             `json:"capacity"`
	AreaSqm     float64         `json:"area_sqm"`
	Floor       int             `json:"floor"`
	Amenities   []string        `json:"amenities"`
	Images      []string        `json:"images"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type RoomShort struct {
	ID          uuid.UUID       `json:"id"`
	HotelID     uuid.UUID       `json:"hotel_id"`
	Description string          `json:"description"`
	RoomNumber  string          `json:"room_number"`
	Type        dto.RoomType    `json:"type"`
	Status      dto.RoomStatus  `json:"status"`
	Price       decimal.Decimal `json:"price"`
	Capacity    int             `json:"capacity"`
	AreaSqm     float64         `json:"area_sqm"`
	Amenities   []string        `json:"amenities"`
	Images      []string        `json:"images"`
}

type RoomUpdate struct {
	Description string          `json:"description"`
	RoomNumber  string          `json:"room_number"`
	Type        dto.RoomType    `json:"type"`
	Price       decimal.Decimal `json:"price"`
	Capacity    int             `json:"capacity"`
	AreaSqm     float64         `json:"area_sqm"`
	Floor       int             `json:"floor"`
	Amenities   []string        `json:"amenities"`
	Images      []string        `json:"images"`
}
type RoomList struct {
	Rooms      []RoomShort `json:"rooms"`
	TotalCount uint64      `json:"total_count"`
}
