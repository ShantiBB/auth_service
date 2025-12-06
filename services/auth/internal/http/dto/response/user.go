package response

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  *string   `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserList struct {
	Users       []User `json:"users"`
	Limit       int    `json:"limit"`
	CurrentPage int    `json:"current_page"`
	HasNextPage bool   `json:"has_next_page"`
	HasPrevPage bool   `json:"has_prev_page"`
	TotalCount  int    `json:"total_count"`
}

type UserShort struct {
	ID       int64   `json:"id"`
	Username *string `json:"username"`
	Email    string  `json:"email"`
}
