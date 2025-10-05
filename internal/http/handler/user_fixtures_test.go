package handler

import (
	"time"

	"auth_service/internal/domain/models"
	"auth_service/internal/http/lib/schemas/request"
)

var userReq = request.UserCreate{
	Email:    "test@example.com",
	Username: "test-user",
	Password: "password123",
}

var userMock = models.User{
	ID:        1,
	Email:     "test@example.com",
	Username:  "test-user",
	Role:      "user",
	IsActive:  true,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
