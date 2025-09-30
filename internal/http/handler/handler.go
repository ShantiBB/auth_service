package handler

import (
	"auth_service/internal/domain/models"
	schemas2 "auth_service/internal/http/schemas"
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

func (h *Handler) UserCreateRequestToEntity(req *schemas2.UserCreateRequest, hash string) *models.UserCreate {
	return &models.UserCreate{
		Username:    req.Username,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Description: req.Description,
		Password:    hash,
	}
}

func (h *Handler) UserEntityToResponse(user *models.User) *schemas2.UserResponse {
	return &schemas2.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Description: user.Description,
	}
}
