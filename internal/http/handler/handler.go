package handler

import (
	"auth_service/internal/domain/models"
	"auth_service/internal/http/schemas"
)

type Service interface {
	UserService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) UserCreateRequestToEntity(req *schemas.UserCreateRequest, hash string) *models.UserCreate {
	return &models.UserCreate{
		Username: req.Username,
		Email:    req.Email,
		Password: hash,
	}
}

func (h *Handler) UserUpdateRequestToEntity(req *schemas.UserUpdateRequest, id int64) *models.User {
	return &models.User{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
	}
}

func (h *Handler) UserEntityToResponse(user *models.User) *schemas.UserResponse {
	return &schemas.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
