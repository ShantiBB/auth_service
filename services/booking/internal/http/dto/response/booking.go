package response

import (
	"booking/internal/repository/models"
	"time"
)

type Booking struct {
	ID          string               `json:"id"`
	UserID      int64                `json:"user_id"`
	HotelID     string               `json:"hotel_id"`
	CheckIn     string               `json:"check_id"`
	CheckOut    string               `json:"check_out"`
	Status      models.BookingStatus `json:"status"`
	GuestName   string               `json:"guest_name"`
	GuestEmail  *string              `json:"guest_email"`
	GuestPhone  *string              `json:"guest_phone"`
	Currency    string               `json:"currency"`
	TotalAmount string               `json:"total_amount"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}
