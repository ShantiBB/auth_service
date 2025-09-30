package schemas

import "time"

type UserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserShortResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
