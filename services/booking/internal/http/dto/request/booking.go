package request

type BookingCreate struct {
	UserID      int64   `json:"user_id"`
	HotelID     string  `json:"hotel_id"`
	CheckIn     string  `json:"check_in"`
	CheckOut    string  `json:"check_out"`
	GuestName   string  `json:"guest_name"`
	GuestEmail  *string `json:"guest_email"`
	GuestPhone  *string `json:"guest_phone"`
	Currency    string  `json:"currency"`
	TotalAmount string  `json:"total_amount"`
}
