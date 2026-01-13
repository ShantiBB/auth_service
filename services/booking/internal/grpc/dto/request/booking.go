package request

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateBooking struct {
	UserID              int64               `json:"user_id"`
	HotelID             string              `json:"hotel_id"`
	CheckIn             time.Time           `json:"check_in"`
	CheckOut            time.Time           `json:"check_out"`
	GuestName           string              `json:"guest_name"`
	GuestEmail          *string             `json:"guest_email"`
	GuestPhone          *string             `json:"guest_phone"`
	Currency            string              `json:"currency"`
	ExpectedTotalAmount decimal.Decimal     `json:"expected_total_amount"`
	Rooms               []CreateBookingRoom `json:"rooms"`
}
