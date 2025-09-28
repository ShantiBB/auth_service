package handler

import (
	"auth_service/api/http/schemas"
	"auth_service/internal/entity"
	"auth_service/internal/service"
)

type Handler struct {
	UserService service.UserService
}

func New(userService service.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) UserCreateRequestToEntity(req *schemas.UserCreateRequest, hash string) *entity.UserCreate {
	return &entity.UserCreate{
		Username:    req.Username,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Description: req.Description,
		Password:    hash,
	}
}

func (h *Handler) UserEntityToResponse(user *entity.User) *schemas.UserResponse {
	return &schemas.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Description: user.Description,
	}
}
