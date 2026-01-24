package models

import "time"

type CreateUser struct {
	Username *string
	Email    string
	Password string
}

type UpdateUser struct {
	ID       int64
	Username string
	Email    string
}

type User struct {
	ID        int64
	Username  *string
	Email     string
	Role      UserRole
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserShort struct {
	ID       int64
	Username *string
	Email    string
	Role     UserRole
	IsActive bool
}

type UserList struct {
	Users      []*UserShort
	TotalCount uint64
}

type UpdateUserPassword struct {
	ID          int64
	Password    string
	NewPassword string
}

type UserCredentials struct {
	ID       int64
	Email    string
	Role     UserRole
	Password string
}

func (u CreateUser) ToUserRead() *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
	}
}
